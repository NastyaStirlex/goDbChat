
-- +migrate Up
alter table incident_service.pretensions
    add is_sent boolean default false not null;

-- +migrate Down
alter table incident_service.pretensions
    drop column is_sent;