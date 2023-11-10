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
    online        BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at    TIMESTAMPTZ          DEFAULT NOW(),
    updated_at    TIMESTAMPTZ          DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER user_updated_at_trigger
    BEFORE UPDATE ON "user"
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TABLE tag
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE user_tag
(
    user_id INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    tag_id  INT NOT NULL REFERENCES tag (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, tag_id)
);

CREATE TABLE "like"
(
    liked_from_user_id INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    liked_to_user_id   INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    created_at         TIMESTAMPTZ DEFAULT NOW(),
    CHECK (liked_from_user_id != liked_to_user_id),
    PRIMARY KEY (liked_from_user_id, liked_to_user_id)
);

CREATE TABLE dialog
(
    id          SERIAL PRIMARY KEY,
    user1_id    INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    user2_id    INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (user1_id, user2_id),
    CONSTRAINT check_pair_order CHECK (user1_id < user2_id)
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
    reporter_user_id  INT      NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    reported_user_id  INT      NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    complaint_type_id INT      NOT NULL REFERENCES complaint_type (id) ON DELETE CASCADE,
    report_status     SMALLINT NOT NULL CHECK (report_status BETWEEN 0 AND 5),
    created_at        TIMESTAMPTZ DEFAULT NOW(),
    CHECK (reporter_user_id != reported_user_id)
);
