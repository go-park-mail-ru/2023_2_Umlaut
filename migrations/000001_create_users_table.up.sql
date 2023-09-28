CREATE TYPE gender AS ENUM (
	'Male',
	'Female'
);

CREATE TABLE users(
    id            serial PRIMARY KEY,
    name          VARCHAR (255) NOT NULL,
    mail          VARCHAR (255) UNIQUE NOT NULL,
    password_hash VARCHAR (255) NOT NULL,
    user_gender   gender,
    prefer_gender gender,
    description   VARCHAR (255),
    age           INT,
    looking       VARCHAR (255),
    education     VARCHAR (255),
    hobbies       VARCHAR (255),
    tags          VARCHAR (255)
);
