-- +migrate Up
create table IF NOT EXISTS users
(
    user_id   int primary key not null,
    name varchar(255)             not null
);

-- +migrate Down
drop table users;