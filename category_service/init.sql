create database category_db;
\c categories_db;
create table categories (
	category_id BIGSERIAL primary key,
	category_name VARCHAR not null,
	created_at TIMESTAMP NOT NULL default current_timestamp,
  	updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
 	deleted_at TIMESTAMP
);
create table book_categories(
	book_category_id BIGSERIAL primary key,
	book_id BIGINT not null,
	category_id BIGINT not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP not null,
	updated_at TIMESTAMP default clock_timestamp() not null,
	deleted_at TIMESTAMP,
	foreign key(category_id) references categories(category_id)
);


INSERT INTO categories (category_name) VALUES ('Fiction');
INSERT INTO categories (category_name) VALUES ('Non-Fiction');
INSERT INTO categories (category_name) VALUES ('Science');
INSERT INTO categories (category_name) VALUES ('History');
INSERT INTO categories (category_name) VALUES ('Biography');
INSERT INTO categories (category_name) VALUES ('Fantasy');
INSERT INTO categories (category_name) VALUES ('Mystery');
INSERT INTO categories (category_name) VALUES ('Romance');
INSERT INTO categories (category_name) VALUES ('Thriller');
INSERT INTO categories (category_name) VALUES ('Self-Help');
INSERT INTO categories (category_name) VALUES ('Health');
INSERT INTO categories (category_name) VALUES ('Travel');
INSERT INTO categories (category_name) VALUES ('Cookbooks');
INSERT INTO categories (category_name) VALUES ('Children');
INSERT INTO categories (category_name) VALUES ('Art');
INSERT INTO categories (category_name) VALUES ('Poetry');
INSERT INTO categories (category_name) VALUES ('Drama');
INSERT INTO categories (category_name) VALUES ('Comics');
INSERT INTO categories (category_name) VALUES ('Religion');
INSERT INTO categories (category_name) VALUES ('Technology');


INSERT INTO book_categories (book_id, category_id) VALUES (1, 1); 
INSERT INTO book_categories (book_id, category_id) VALUES (2, 2); 
INSERT INTO book_categories (book_id, category_id) VALUES (3, 3); 
INSERT INTO book_categories (book_id, category_id) VALUES (4, 4); 
INSERT INTO book_categories (book_id, category_id) VALUES (5, 5); 
INSERT INTO book_categories (book_id, category_id) VALUES (6, 6); 
INSERT INTO book_categories (book_id, category_id) VALUES (7, 7); 
INSERT INTO book_categories (book_id, category_id) VALUES (8, 8); 
INSERT INTO book_categories (book_id, category_id) VALUES (9, 9);
INSERT INTO book_categories (book_id, category_id) VALUES (10, 10); 
INSERT INTO book_categories (book_id, category_id) VALUES (1, 11); 
INSERT INTO book_categories (book_id, category_id) VALUES (2, 12); 
INSERT INTO book_categories (book_id, category_id) VALUES (3, 13); 
INSERT INTO book_categories (book_id, category_id) VALUES (4, 14); 
INSERT INTO book_categories (book_id, category_id) VALUES (5, 15); 
INSERT INTO book_categories (book_id, category_id) VALUES (6, 16); 
INSERT INTO book_categories (book_id, category_id) VALUES (7, 17); 
INSERT INTO book_categories (book_id, category_id) VALUES (8, 18); 
INSERT INTO book_categories (book_id, category_id) VALUES (9, 19); 
INSERT INTO book_categories (book_id, category_id) VALUES (10, 20); 