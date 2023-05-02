-- +migrate Up
create table IF NOT EXISTS results
(
    mem_id int not null,
    user_id   int not null,
    isLike bool,
    primary key (mem_id, user_id)
);

-- +migrate Down
drop table results;