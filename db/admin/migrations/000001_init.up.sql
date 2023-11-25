CREATE TABLE statistic
(
    id          SERIAL PRIMARY KEY,
    user_id     INT NOT NULL,
    rating      INT,
    liked       TEXT,
    need_fix    TEXT,
    comment_fix TEXT,
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
