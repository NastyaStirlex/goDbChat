-- +migrate Up
alter table mems_db.public.results
    add constraint fk_mems_results foreign key (mem_id) references mems (mem_id);

alter table mems_db.public.results
    add constraint fk_users_results foreign key (user_id) references users (user_id);

-- +migrate Down
alter table mems_db.public.results
    drop constraint fk_mems_results;

alter table mems_db.public.results
    drop constraint fk_users_results;