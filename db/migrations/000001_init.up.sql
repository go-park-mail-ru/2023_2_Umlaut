CREATE TABLE "user"
(
    id            SERIAL PRIMARY KEY,
    name          TEXT NOT NULL,
    gender        SMALLINT,
    prefer_gender SMALLINT,
    description   TEXT,
    interests     TEXT,
    image_path    TEXT,
    education     TEXT,
    hobbies       TEXT,
    birthday      DATE,
    online        BOOLEAN DEFAULT FALSE
);

CREATE TABLE credential
(
    user_id       SERIAL PRIMARY KEY REFERENCES "user" (id) ON DELETE CASCADE,
    mail          TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    salt          TEXT NOT NULL
);

CREATE TABLE tag
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE user_tag
(
    user_id INT REFERENCES "user" (id) ON DELETE CASCADE,
    tag_id  INT REFERENCES tag (id) ON DELETE CASCADE
);

CREATE TABLE "like"
(
    liked_by_user_id INT REFERENCES "user" (id) ON DELETE CASCADE,
    liked_to_user_id INT REFERENCES "user" (id) ON DELETE CASCADE,
    committed_at     TIMESTAMP DEFAULT NOW()
);

CREATE TABLE dialog
(
    id       SERIAL PRIMARY KEY,
    user1_id INT REFERENCES "user" (id) ON DELETE SET NULL,
    user2_id INT REFERENCES "user" (id) ON DELETE SET NULL
);

CREATE TABLE message
(
    id           SERIAL PRIMARY KEY,
    dialog_id    INT REFERENCES dialog (id) ON DELETE CASCADE,
    sender_id    INT REFERENCES "user" (id) ON DELETE SET NULL,
    message_text TEXT NOT NULL,
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
    reporter_user_id  INT REFERENCES "user" (id) ON DELETE CASCADE,
    reported_user_id  INT REFERENCES "user" (id) ON DELETE CASCADE,
    complaint_type_id INT REFERENCES complaint_type (id),
    report_status     SMALLINT NOT NULL,
    complaint_time    TIMESTAMP DEFAULT NOW()
);
