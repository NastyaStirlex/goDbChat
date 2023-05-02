-- +migrate Up
create schema IF NOT EXISTS mems_db;

-- +migrate Down
drop schema IF EXISTS mems_db;