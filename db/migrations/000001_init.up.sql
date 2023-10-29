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
    committed_at     TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE dialog
(
    id       SERIAL PRIMARY KEY,
    user1_id INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    user2_id INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL
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
