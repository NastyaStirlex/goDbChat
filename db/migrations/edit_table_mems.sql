-- +migrate Up
alter table mems_db.public.mems
    ALTER COLUMN mem_link TYPE VARCHAR;