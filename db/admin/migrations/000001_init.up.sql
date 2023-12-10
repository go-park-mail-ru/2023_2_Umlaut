CREATE TABLE feedback
(
    id          SERIAL PRIMARY KEY,
    user_id     INT NOT NULL,
    rating      INT,
    liked       TEXT,
    need_fix    TEXT,
    comment     TEXT,
    created_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE recommendation
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    recommend  INT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE feed_feedback
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    recommend  INT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE admin
(
    id            SERIAL PRIMARY KEY,
    name          TEXT        NOT NULL,
    mail          TEXT UNIQUE NOT NULL,
    password_hash TEXT        NOT NULL,
    salt          TEXT        NOT NULL
);

INSERT INTO admin (name, mail, password_hash, salt)
VALUES ('admin', 'admin@admin.ru',
        '635262426a51506543766a5078476349596d747150577c4a8d09ca3762af61e59520943dc26494f8941b',
        'cRbBjQPeCvjPxGcIYmtqPW')