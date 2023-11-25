CREATE TABLE feedback
(
    id          SERIAL PRIMARY KEY,
    user_id     INT NOT NULL,
    rating      INT,
    liked       TEXT,
    need_fix    TEXT,
    comment_fix TEXT,
    comment     TEXT,
    show        BOOlEAN,
    created_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE recommendation
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    recommend  INT,
    show       BOOlEAN,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE feed_feedback
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL,
    recommend  INT,
    show       BOOlEAN,
    created_at TIMESTAMP DEFAULT NOW()
);
