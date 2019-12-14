create table if not exists server(
    id                      varchar(255) not null,
    name                    varchar(255) not null,
    is_connected            boolean default false,
    wrong_protocol          boolean default false,
    organization_name       varchar(100) not null,
    organization_fullname   varchar(255) not null,
    organization_phone      varchar(20),
    organization_fax        varchar(100),
    organization_email      varchar(100)
);