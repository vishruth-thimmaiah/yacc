CREATE TABLE users (
    id uuid primary key default(uuid_generate_v4()),
	username text,
	email text not null unique,
	passwd char(65) not null
);

CREATE TABLE session (
	sessionid uuid primary key default(uuid_generate_v4()),
    id uuid not null references users(id) on delete cascade
);

CREATE TABLE messages (
	senderid uuid not null references users(id) on delete cascade,
	chat_id uuid not null references contacts(chat_id) on delete cascade,
	message text not null,
	date timestamp with time zone default(now()),
	attachment text
);

CREATE TABLE contacts (
    user1 uuid not null references users(id) on delete cascade,
    user2 uuid not null references users(id) on delete cascade,
    chat_id uuid default(uuid_generate_v4()) primary key
);
