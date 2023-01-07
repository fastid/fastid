create table users
(
    user_id uuid not null primary key,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    username varchar(200) null,
    email varchar(200) null,
    password char(92) null,
    is_active boolean not null default 'false',
    is_superuser boolean not null default 'false'
);


create index users_username_index on users (username);
create index users_email_index on users (email);
