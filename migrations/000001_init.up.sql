CREATE TABLE users (
    id            serial PRIMARY KEY,
    name          VARCHAR (255) NOT NULL,
    mail          VARCHAR (255) UNIQUE NOT NULL,
    password_hash VARCHAR (255) NOT NULL,
    salt          VARCHAR (255) NOT NULL,
    user_gender   VARCHAR (255),
    prefer_gender VARCHAR (255),
    description   VARCHAR (255),
    age           INT,
    looking       VARCHAR (255),
    education     VARCHAR (255),
    hobbies       VARCHAR (255),
    tags          VARCHAR (255)
);
