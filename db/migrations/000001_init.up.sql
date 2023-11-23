CREATE TABLE "user"
(
    id            SERIAL PRIMARY KEY,
    name          TEXT        NOT NULL,
    mail          TEXT UNIQUE NOT NULL,
    password_hash TEXT        NOT NULL,
    salt          TEXT        NOT NULL,
    user_gender   SMALLINT CHECK (user_gender BETWEEN 0 AND 1),
    prefer_gender SMALLINT CHECK (prefer_gender BETWEEN 0 AND 1),
    description   TEXT,
    looking       TEXT,
    image_paths   TEXT[]               DEFAULT ARRAY []::TEXT[],
    education     TEXT,
    hobbies       TEXT,
    birthday      DATE,
    online        BOOLEAN     NOT NULL DEFAULT FALSE,
    tags          TEXT[]               DEFAULT ARRAY []::TEXT[],
    created_at    TIMESTAMPTZ          DEFAULT NOW(),
    updated_at    TIMESTAMPTZ          DEFAULT NOW()
);

CREATE TABLE tag
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE "like"
(
    liked_by_user_id INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    liked_to_user_id INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    created_at       TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (liked_by_user_id, liked_to_user_id)
);

CREATE TABLE dialog
(
    id         SERIAL PRIMARY KEY,
    user1_id   INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    user2_id   INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
--     last_message_id INT REFERENCES message (id) ON DELETE SET NULL DEFAULT NULL, не раскоментирывать!
    UNIQUE (user1_id, user2_id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT check_pair_order CHECK (user1_id < user2_id)
);

CREATE TABLE message
(
    id           SERIAL PRIMARY KEY,
    dialog_id    INT  NOT NULL REFERENCES dialog (id) ON DELETE CASCADE,
    sender_id    INT  NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    message_text TEXT NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT NOW()
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
    reporter_user_id  INT      NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    reported_user_id  INT      NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    complaint_type_id INT      NOT NULL REFERENCES complaint_type (id) ON DELETE CASCADE,
    report_status     SMALLINT NOT NULL CHECK (report_status BETWEEN 0 AND 5),
    created_at        TIMESTAMPTZ DEFAULT NOW(),
    CHECK (reporter_user_id != reported_user_id)
);


-- triggers
CREATE OR REPLACE FUNCTION delete_tag_cascade()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE "user" SET tags = array_remove(tags, OLD.name) WHERE OLD.name = ANY (tags);
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_delete_tag_cascade
    AFTER DELETE
    ON tag
    FOR EACH ROW
EXECUTE FUNCTION delete_tag_cascade();

CREATE OR REPLACE FUNCTION update_last_message_id()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE dialog SET last_message_id = NEW.id WHERE id = NEW.dialog_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_last_message_id
    AFTER INSERT
    ON message
    FOR EACH ROW
EXECUTE FUNCTION update_last_message_id();


CREATE OR REPLACE FUNCTION delete_invalid_tag()
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
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_delete_invalid_tag
    BEFORE INSERT or UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE FUNCTION delete_invalid_tag();


CREATE OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER user_updated_at_trigger
    BEFORE UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();


-- fill db
INSERT INTO tag (name)
VALUES ('Спорт'),
       ('Музыка'),
       ('Путешествия'),
       ('Еда'),
       ('Искусство'),
       ('Наука');

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

INSERT INTO dialog (user1_id, user2_id)
VALUES (3, 4),
       (3, 5),
       (3, 6),
       (3, 7),
       (3, 8),
       (3, 9);

INSERT INTO message (dialog_id, sender_id, message_text)
VALUES (2, 3, 'Hello'),
       (2, 4, 'Hello last'),
       (3, 5, 'Hello last 1');
