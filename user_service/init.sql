	create database user_service_db;
    \c user_service_db;

	create type gender_enum as enum ('male','female');
	create type role_enum as enum ('admin','user');
	create table users(
		user_id BIGSERIAL not null,
		user_name VARCHAR not null,
		email VARCHAR not null,
		user_password VARCHAR,
		date_of_birth TIMESTAMP not null,
		gender gender_enum not null,
		user_role  role_enum not null,
		photo_url VARCHAR not null,
		created_at TIMESTAMP not null DEFAULT(NOW()),
		updated_at TIMESTAMP not null default CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP
	);
	
	create type token_enum as enum ('verification', 'reset','refresh');
	create table tokens(
		token_id BIGSERIAL not null,
		type token_enum not null,
		user_id BIGINT not null,
		expired_at TIMESTAMP not null,
		created_at TIMESTAMP not null DEFAULT(NOW()),
		updated_at TIMESTAMP not null default CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP,
		foreign Key(user_id) references users(user_id)
	)


INSERT INTO users
(user_name,user_role, date_of_birth, email, user_password, gender, photo_url, address, created_at, updated_at, deleted_at)
VALUES('John Doe','user', '1999-06-04', 'john.doe@mail.com', '$2a$12$ImKhfp3w9TuCtWz2VlxisOPX0iQictBFN6o.QGJSkQmK5FQrFZxBq', 'male'::public.gender_type, '', 'sumatra cina', '2024-08-04 16:59:17.232', '2024-08-04 16:59:17.232', NULL);
INSERT INTO users
(user_name, date_of_birth, email, user_password, gender, photo_url, address, created_at, updated_at, deleted_at)
VALUES('John Doe','user', '1999-06-04', 'john.dgoe@mail.com', '$2a$12$gGZgD9US4P.RfDOeBHlkluFLeO4rxpwNTccLULubpEqch4Mim77la', 'male'::public.gender_type, '', '', '2024-08-04 16:55:13.714', '2024-08-04 16:55:13.715', NULL);
INSERT INTO users
(user_name, date_of_birth, email, user_password, gender, photo_url, address, created_at, updated_at, deleted_at)
VALUES('John Di','admin', '1999-06-04', 'admin@mail.com', '$2a$12$gGZgD9US4P.RfDOeBHlkluFLeO4rxpwNTccLULubpEqch4Mim77la', 'male'::public.gender_type, '', '', '2024-08-04 16:55:13.714', '2024-08-04 16:55:13.715', NULL);
