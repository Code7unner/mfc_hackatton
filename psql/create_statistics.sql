create table if not exists statistics(
    id                          serial not null primary key,
    average_awaiting_time       integer not null,
    active_work_places_count    integer default 0,
    completed_tickets_count     integer not null,
    pending_tikets_count        integer not null
);