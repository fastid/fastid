create table keys
(
    unpacking_key  bytea    not null,
    signature_key  bytea    not null
);

-- create table users
-- (
--     user_id    uuid                                               not null primary key,
--     created_at timestamp with time zone default CURRENT_TIMESTAMP not null,
--     updated_at timestamp with time zone default CURRENT_TIMESTAMP not null,
--     sub        bytea                                              not null,
--     hash_sub   char(40)                                           not null constraint users_hash_sub unique
-- );
