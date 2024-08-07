CREATE DATABASE user_service_db;
\c user_service_db;

CREATE TYPE gender_enum AS ENUM ('male', 'female');
CREATE TYPE role_enum AS ENUM ('admin', 'user');
CREATE TYPE token_enum AS ENUM ('verification', 'reset', 'refresh');

CREATE TABLE users (
    user_id BIGSERIAL NOT NULL PRIMARY KEY,
    user_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    user_password VARCHAR,
    date_of_birth TIMESTAMP NOT NULL,
    gender gender_enum NOT NULL,
    user_role role_enum NOT NULL,
    photo_url VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE tokens (
    token_id BIGSERIAL NOT NULL PRIMARY KEY,
    type token_enum NOT NULL,
    user_id BIGINT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

INSERT INTO users (user_name, user_role, date_of_birth, email, user_password, gender, photo_url, created_at, updated_at, deleted_at)
VALUES ('John Doe', 'user', '1999-06-04', 'john.doe@mail.com', '$2a$12$ImKhfp3w9TuCtWz2VlxisOPX0iQictBFN6o.QGJSkQmK5FQrFZxBq', 'male'::gender_enum, '', '2024-08-04 16:59:17.232', '2024-08-04 16:59:17.232', NULL);

INSERT INTO users (user_name, user_role, date_of_birth, email, user_password, gender, photo_url, created_at, updated_at, deleted_at)
VALUES ('John Doe', 'user', '1999-06-04', 'john.dgoe@mail.com', '$2a$12$gGZgD9US4P.RfDOeBHlkluFLeO4rxpwNTccLULubpEqch4Mim77la', 'male'::gender_enum, '', '2024-08-04 16:55:13.714', '2024-08-04 16:55:13.715', NULL);

INSERT INTO users (user_name, user_role, date_of_birth, email, user_password, gender, photo_url, created_at, updated_at, deleted_at)
VALUES ('John Di', 'admin', '1999-06-04', 'admin@mail.com', '$2a$12$gGZgD9US4P.RfDOeBHlkluFLeO4rxpwNTccLULubpEqch4Mim77la', 'male'::gender_enum, '', '2024-08-04 16:55:13.714', '2024-08-04 16:55:13.715', NULL);
