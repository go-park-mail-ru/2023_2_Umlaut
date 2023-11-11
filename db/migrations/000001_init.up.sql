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
    image_path    TEXT,
    education     TEXT,
    birthday      DATE        NOT NULL,
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
    liked_from_user_id INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    liked_to_user_id   INT NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
    created_at         TIMESTAMPTZ DEFAULT NOW(),
    CHECK (liked_from_user_id != liked_to_user_id),
    PRIMARY KEY (liked_from_user_id, liked_to_user_id)
);

CREATE TABLE dialog
(
    id              SERIAL PRIMARY KEY,
    user1_id        INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    user2_id        INT NOT NULL REFERENCES "user" (id) ON DELETE SET NULL,
    created_at      TIMESTAMPTZ                                    DEFAULT NOW(),
    last_message_id INT REFERENCES message (id) ON DELETE SET NULL DEFAULT NULL,
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


CREATE OR REPLACE FUNCTION delete_invalid_tag()
    RETURNS TRIGGER AS
$$
BEGIN
    FOR new_tag IN NEW.tags
        LOOP
            IF new_tag NOT IN (SELECT name FROM tag) THEN
                NEW.tsgs = array_remove(NEW.tsgs, new_tag);
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
