
DROP DATABASE IF EXISTS author_db;

CREATE DATABASE author_db;

\c author_db;

CREATE TYPE author_gender AS ENUM ('male', 'female', 'unknown');

CREATE TABLE authors (
    author_id BIGSERIAL PRIMARY KEY,
    author_name VARCHAR NOT NULL,
    photo_url VARCHAR NOT NULL,
    gender author_gender NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

create table author_books(
	book_author_id BIGSERIAL primary key,
	book_id BIGINT not null,
	author_id BIGINT not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP not null,
	updated_at TIMESTAMP default clock_timestamp() not null,
	deleted_at TIMESTAMP,
	foreign Key(author_id) references authors(author_id)
);
INSERT INTO authors (author_name, photo_url, gender) VALUES
('Author 1', 'https://example.com/photo1.jpg', 'male'),
('Author 2', 'https://example.com/photo2.jpg', 'female'),
('Author 3', 'https://example.com/photo3.jpg', 'male'),
('Author 4', 'https://example.com/photo4.jpg', 'female'),
('Author 5', 'https://example.com/photo5.jpg', 'male'),
('Author 6', 'https://example.com/photo6.jpg', 'female'),
('Author 7', 'https://example.com/photo7.jpg', 'male'),
('Author 8', 'https://example.com/photo8.jpg', 'female'),
('Author 9', 'https://example.com/photo9.jpg', 'male'),
('Author 10', 'https://example.com/photo10.jpg', 'female'),
('Author 11', 'https://example.com/photo11.jpg', 'male'),
('Author 12', 'https://example.com/photo12.jpg', 'female'),
('Author 13', 'https://example.com/photo13.jpg', 'male'),
('Author 14', 'https://example.com/photo14.jpg', 'female'),
('Author 15', 'https://example.com/photo15.jpg', 'male'),
('Author 16', 'https://example.com/photo16.jpg', 'female'),
('Author 17', 'https://example.com/photo17.jpg', 'male'),
('Author 18', 'https://example.com/photo18.jpg', 'female'),
('Author 19', 'https://example.com/photo19.jpg', 'male'),
('Author 20', 'https://example.com/photo20.jpg', 'female');

INSERT INTO author_books (book_id, author_id) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 2),
(5, 3),
(6, 3),
(7, 4),
(8, 4),
(9, 5),
(10, 5),
(11, 6),
(12, 6),
(13, 7),
(14, 7),
(15, 8),
(16, 8),
(17, 9),
(18, 9),
(19, 10),
(20, 10);
