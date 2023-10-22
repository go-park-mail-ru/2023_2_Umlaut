CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    user_gender SMALLINT,
    prefer_gender SMALLINT,
    description TEXT,
    interests TEXT,
    image_path TEXT,
    education TEXT,
    hobbies TEXT,
    birthday DATE,
    online BOOLEAN
);

CREATE TABLE credentials (
    user_id SERIAL PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
    mail TEXT,
    password_hash TEXT,
    salt TEXT
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE user_tags (
    user_id INT REFERENCES users (id) ON DELETE CASCADE,
    tag_id INT REFERENCES tags (id) ON DELETE CASCADE
);

CREATE TABLE likes (
    liked_by_user_id INT REFERENCES users (id) ON DELETE CASCADE,
    liked_to_user_id INT REFERENCES users (id) ON DELETE CASCADE,
    committed_at TIMESTAMP
);

CREATE TABLE dialogs (
    id SERIAL PRIMARY KEY,
    user1_id INT REFERENCES users (id) ON DELETE SET NULL,
    user2_id INT REFERENCES users (id) ON DELETE SET NULL
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    dialog_id INT REFERENCES dialogs (id) ON DELETE CASCADE,
    sender_id INT REFERENCES users (id) ON DELETE SET NULL,
    message_text TEXT,
    message_time TIMESTAMP
);

CREATE TABLE complaint_types (
    id SERIAL PRIMARY KEY,
    type_name TEXT
);

CREATE TABLE complaints (
    id SERIAL PRIMARY KEY,
    reporter_user_id INT REFERENCES users (id) ON DELETE CASCADE,
    reported_user_id INT REFERENCES users (id) ON DELETE CASCADE,
    complaint_type_id INT REFERENCES complaint_types (id),
    report_status SMALLINT,
    complaint_time TIMESTAMP
);
