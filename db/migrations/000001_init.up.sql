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
    image_path    TEXT,
    education     TEXT,
    hobbies       TEXT,
    birthday      DATE,
    online        BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE TABLE tag
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE user_tag
(
    user_id INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    tag_id  INT NOT NULL REFERENCES tag (id) ON DELETE CASCADE
);

CREATE TABLE "like"
(
    liked_by_user_id INT       NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    liked_to_user_id INT       NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    committed_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    liked_to_user_id INT       NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    UNIQUE (liked_by_user_id, liked_to_user_id)
);

CREATE TABLE dialog
(
    id       SERIAL PRIMARY KEY,
    user1_id INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    user2_id INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    UNIQUE (user1_id, user2_id)
);

CREATE TABLE message
(
    id           SERIAL PRIMARY KEY,
    dialog_id    INT       NOT NULL REFERENCES dialog (id) ON DELETE CASCADE,
    sender_id    INT       NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    message_text TEXT      NOT NULL,
    message_time TIMESTAMP NOT NULL
);

CREATE TABLE complaint_type
(
    id        SERIAL PRIMARY KEY,
    type_name TEXT UNIQUE NOT NULL
);

CREATE TABLE complaint
(
    id                SERIAL PRIMARY KEY,
    reporter_user_id  INT       NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    reported_user_id  INT       NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    complaint_type_id INT       NOT NULL REFERENCES complaint_type (id) ON DELETE CASCADE,
    report_status     SMALLINT  NOT NULL CHECK (report_status BETWEEN 0 AND 5),
    complaint_time    TIMESTAMP NOT NULL DEFAULT NOW()
);

-- fill db

INSERT INTO "user" (name, mail, password_hash, salt, user_gender, prefer_gender, description, looking, image_path, education, hobbies, birthday)
VALUES
    ('Фёдор', 'fedor@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент второго семестра технопарка, backend', 'Новые знакомства', NULL, 'Неполное высшее', 'Баскетбол, сноуборд', '2003-01-01'),
    ('Дмитрий', 'dmitry@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент второго семестра технопарка, frontend', 'Новые знакомства', NULL, 'Неполное высшее', 'Волейбол, авиамоделирование', '2003-01-01'),
    ('Максим', 'max@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент второго семестра технопарка, backend', 'Новые знакомства', NULL, 'Неполное высшее', 'Футбол, керлинг', '2003-01-01'),
    ('Иван', 'ivan@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент 3 курса МГТУ им. Н. Э. Баумана', 'Серьезные отношения', NULL, 'Неполное высшее', 'Волейбол, футбол', '2003-01-01'),
    ('Алексей', 'Alexey@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 1, 0, 'Студент 2 курса МФТИ', 'Серьезные отношения', NULL, 'Неполное высшее', 'Компьютерные игры', '2003-01-01'),
    ('Полина', 'polina@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Студент второго семестра технопарка, frontend', 'Новые знакомства', NULL, 'Неполное высшее', 'Баскетбол, сноуборд', '2003-01-01'),
    ('Ирина', 'irina@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Учусь в школе, собираюсь поступать в МГТУ', 'Серьезные отношения', NULL, 'Неполное среднее', 'Дзюдо, бокс', '2005-01-01'),
    ('Анна', 'anna@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Студент МГТУ, направление 09.03.03', 'Новые знакомства', NULL, 'Неполное высшее', 'Танцы, настольные игры', '2003-01-01'),
    ('Карина', 'karina@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Спортшкольница, играю в сборной по футболу', 'Серьезные отношения', NULL, 'Полное высшее', 'Футбол, фехтование', '2001-01-01'),
    ('Юлия', 'julia@mail.ru', '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b', 'cRbBjQPeCvjPxGcIYmtqPW', 0, 1, 'Студент 1 курса МГУ', 'Новые знакомства', NULL, 'Неполное высшее', 'Волейбол, компьютерные игры', '2003-01-01');
