CREATE TABLE users (
	uuid bigserial not null primary key,
	username text,
	email text not null unique,
	passwd char(66) not null,
);
