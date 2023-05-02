-- +migrate Up
create table IF NOT EXISTS mems
(
    mem_id   int primary key not null,
    mem_link int             not null,
    likes int             not null,
    dislikes int not null
);

-- +migrate Down
drop table mems;