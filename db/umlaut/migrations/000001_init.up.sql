CREATE
    OR REPLACE FUNCTION calculate_age(birth_date DATE)
    RETURNS INTEGER AS
$$
BEGIN
    RETURN DATE_PART('year', CURRENT_DATE) - DATE_PART('year', birth_date);
END;
$$
    LANGUAGE plpgsql
    IMMUTABLE;

CREATE TABLE "user"
(
    id            SERIAL PRIMARY KEY,
    name          TEXT     NOT NULL,
    mail          TEXT UNIQUE,
    password_hash TEXT,
    salt          TEXT,
    user_gender   SMALLINT CHECK (user_gender BETWEEN 0 AND 1),
    prefer_gender SMALLINT CHECK (prefer_gender BETWEEN 0 AND 1),
    description   TEXT,
    looking       TEXT,
    image_paths   TEXT[]            DEFAULT ARRAY []::TEXT[],
    education     TEXT,
    hobbies       TEXT,
    birthday      DATE,
    role          SMALLINT NOT NULL DEFAULT 1 CHECK (role BETWEEN 1 AND 3),
    invited_by    INT      REFERENCES "user" (id) ON DELETE SET NULL,
    like_counter  INT               DEFAULT 30,
    online        BOOLEAN  NOT NULL DEFAULT FALSE,
    tags          TEXT[]            DEFAULT ARRAY []::TEXT[],
    age           INTEGER GENERATED ALWAYS AS (calculate_age(birthday)) STORED,
    oauth_id      INT UNIQUE,
    created_at    TIMESTAMPTZ       DEFAULT timezone('Europe/Moscow'::text, NOW()),
    updated_at    TIMESTAMPTZ       DEFAULT timezone('Europe/Moscow'::text, NOW())
);

CREATE TABLE tag
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE "like"
(
    liked_by_user_id INT     NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    liked_to_user_id INT     NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    is_like          BOOLEAN NOT NULL,
    created_at       TIMESTAMPTZ DEFAULT timezone('Europe/Moscow'::text, NOW()),
    UNIQUE (liked_by_user_id, liked_to_user_id)
);

CREATE TABLE dialog
(
    id         SERIAL PRIMARY KEY,
    user1_id   INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    user2_id   INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    banned     BOOlEAN     DEFAULT FALSE,
--     last_message_id INT REFERENCES message (id) ON DELETE SET NULL DEFAULT NULL, не раскоментирывать!
    UNIQUE (user1_id, user2_id),
    created_at TIMESTAMPTZ DEFAULT timezone('Europe/Moscow'::text, NOW()),
    CONSTRAINT check_pair_order CHECK (user1_id < user2_id)
);

CREATE TABLE message
(
    id           SERIAL PRIMARY KEY,
    dialog_id    INT  NOT NULL REFERENCES dialog (id) ON DELETE CASCADE,
    sender_id    INT  NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    recipient_id INT  NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    message_text TEXT NOT NULL,
    is_read      BOOLEAN     DEFAULT FALSE,
    created_at   TIMESTAMPTZ DEFAULT timezone('Europe/Moscow'::text, NOW())
);

ALTER TABLE dialog
    ADD last_message_id INT REFERENCES message (id) ON DELETE SET NULL DEFAULT NULL;

CREATE TABLE complaint_type
(
    id        SERIAL PRIMARY KEY,
    type_name TEXT UNIQUE NOT NULL
);

CREATE TABLE complaint
(
    id                SERIAL PRIMARY KEY,
    reporter_user_id  INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    reported_user_id  INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    complaint_type_id INT NOT NULL REFERENCES "complaint_type" (id) ON DELETE CASCADE,
    complaint_text    TEXT,
    report_status     SMALLINT    DEFAULT 0,
    created_at        TIMESTAMPTZ DEFAULT timezone('Europe/Moscow'::text, NOW()),
    UNIQUE (reporter_user_id, reported_user_id),
    CHECK (reporter_user_id != reported_user_id)
);


-- triggers
CREATE
    OR REPLACE FUNCTION delete_tag_cascade()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE "user"
    SET tags = array_remove(tags, OLD.name)
    WHERE OLD.name = ANY (tags);
    RETURN OLD;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trigger_delete_tag_cascade
    AFTER DELETE
    ON tag
    FOR EACH ROW
EXECUTE FUNCTION delete_tag_cascade();

CREATE
    OR REPLACE FUNCTION update_last_message_id()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE dialog
    SET last_message_id = NEW.id
    WHERE id = NEW.dialog_id;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_last_message_id
    AFTER INSERT
    ON message
    FOR EACH ROW
EXECUTE FUNCTION update_last_message_id();


CREATE
    OR REPLACE FUNCTION delete_invalid_tag()
    RETURNS TRIGGER AS
$$
DECLARE
    new_tag text;
BEGIN
    FOREACH new_tag IN ARRAY NEW.tags
        LOOP
            IF new_tag NOT IN (SELECT name FROM tag) THEN
                NEW.tags = array_remove(NEW.tags, new_tag);
            END IF;
        END LOOP;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trigger_delete_invalid_tag
    BEFORE INSERT or
        UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE FUNCTION delete_invalid_tag();


CREATE
    OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = timezone('Europe/Moscow'::text, NOW());
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER user_updated_at_trigger
    BEFORE UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE
    OR REPLACE FUNCTION update_banned_dialog()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE dialog
    SET banned = TRUE
    WHERE user1_id = LEAST(NEW.reporter_user_id, NEW.reported_user_id)
      AND user2_id = GREATEST(NEW.reporter_user_id, NEW.reported_user_id);
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_banned_dialog
    AFTER INSERT
    ON complaint
    FOR EACH ROW
EXECUTE FUNCTION update_banned_dialog();

CREATE
    OR REPLACE FUNCTION update_banned_user()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE "user"
    SET role = 3
    WHERE id = NEW.reported_user_id;
    UPDATE dialog
    SET banned = TRUE
    WHERE user1_id = NEW.reported_user_id
       OR user2_id = NEW.reported_user_id;
    DELETE
    FROM complaint
    WHERE reported_user_id = NEW.reported_user_id
      AND id != NEW.id;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_banned_user
    AFTER UPDATE
    ON complaint
    FOR EACH ROW
EXECUTE FUNCTION update_banned_user();


CREATE
    OR REPLACE FUNCTION update_user_role()
    RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.invited_by IS NOT NULL AND
       (SELECT count(id)
        FROM "user"
        WHERE invited_by = NEW.invited_by
          AND description IS NOT NULL) = 5
    THEN
        UPDATE "user"
        SET role         = 2,
            like_counter = -1
        WHERE id = NEW.invited_by;
    END IF;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_user_role
    AFTER UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE FUNCTION update_user_role();


CREATE
    OR REPLACE FUNCTION update_user_like_counter()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE "user"
    SET like_counter = like_counter - 1
    WHERE id = NEW.liked_by_user_id
      AND role != 2;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_user_like_counter
    AFTER INSERT
    ON "like"
    FOR EACH ROW
EXECUTE FUNCTION update_user_like_counter();


-- fill db
INSERT INTO tag (name)
VALUES ('Романтика'),
       ('Путешествия'),
       ('Фитнес'),
       ('Кулинария'),
       ('Искусство'),
       ('Музыка'),
       ('Фотография'),
       ('Литература'),
       ('Технологии'),
       ('Экология'),
       ('Кино и телевидение'),
       ('Спорт'),
       ('Психология'),
       ('Домашние животные'),
       ('Игры'),
       ('Автомобили'),
       ('Финансы и бизнес'),
       ('Мода'),
       ('Природа'),
       ('Образование');

INSERT INTO "user" (name, mail, password_hash, salt, user_gender, prefer_gender, description, looking, image_paths,
                    education, hobbies, birthday, tags)
VALUES ('Фёдор', 'fedor@mail.ru',
        '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент второго семестра технопарка, backend', 'Новые знакомства', NULL,
        'Неполное высшее', 'Баскетбол, сноуборд', '2003-01-01', ARRAY ['Спорт', 'Музыка', 'Путешествия']),
       ('Дмитрий', 'dmitry@mail.ru',
        '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент второго семестра технопарка, frontend', 'Новые знакомства', NULL,
        'Неполное высшее', 'Волейбол, авиамоделирование', '2003-01-01', ARRAY ['Спорт', 'Музыка', 'Путешествия']),
       ('Максим', 'max@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент второго семестра технопарка, backend', 'Новые знакомства', NULL,
        'Неполное высшее', 'Футбол, керлинг', '2003-01-01', ARRAY ['Спорт', 'Музыка', 'Путешествия']),
       ('Иван', 'ivan@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент 3 курса МГТУ им. Н. Э. Баумана', 'Серьезные отношения', NULL,
        'Неполное высшее', 'Волейбол, футбол', '2003-01-01', ARRAY ['Музыка', 'Еда', 'Искусство']),
       ('Алексей', 'Alexey@mail.ru',
        '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент 2 курса МФТИ', 'Серьезные отношения', NULL, 'Неполное высшее',
        'Компьютерные игры', '2003-01-01', ARRAY ['Музыка', 'Еда', 'Искусство']),
       ('Полина', 'polina@mail.ru',
        '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Студент второго семестра технопарка, frontend', 'Новые знакомства', NULL,
        'Неполное высшее', 'Баскетбол, сноуборд', '2003-01-01', ARRAY ['Музыка', 'Еда', 'Искусство']),
       ('Ирина', 'irina@mail.ru',
        '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Учусь в школе, собираюсь поступать в МГТУ', 'Серьезные отношения', NULL,
        'Неполное среднее', 'Дзюдо, бокс', '2005-01-01', ARRAY ['Музыка', 'Еда', 'Искусство']),
       ('Анна', 'anna@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Студент МГТУ, направление 09.03.03', 'Новые знакомства', NULL,
        'Неполное высшее', 'Танцы, настольные игры', '2003-01-01', ARRAY ['Музыка', 'Еда', 'Искусство']),
       ('Карина', 'karina@mail.ru',
        '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Спортшкольница, играю в сборной по футболу', 'Серьезные отношения', NULL,
        'Полное высшее', 'Футбол, фехтование', '2001-01-01', ARRAY ['Путешествия', 'Наука']),
       ('Юлия', 'julia@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Студент 1 курса МГУ', 'Новые знакомства', NULL, 'Неполное высшее',
        'Волейбол, компьютерные игры', '2003-01-01', ARRAY ['Путешествия', 'Наука']);

INSERT INTO complaint_type (type_name)
VALUES ('Порнография'),
       ('Рассылка спама'),
       ('Оскорбительное поведение'),
       ('Мошенничество'),
       ('Рекламная страница'),
       ('Клон моей страницы (или моя старая страница)'),
       ('Другое');
