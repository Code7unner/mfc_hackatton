create table if not exists statistics(
    id                          varchar(255) not null,
    average_awaiting_time       integer not null,
    active_work_places_count    integer default 0,
    completed_tickets_count     integer not null,
    pending_tikets_count        integer not null
);