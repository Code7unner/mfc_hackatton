create table if not exists servers(
    id text not null,
    name text not null,
    is_connected boolean default false,
    wrong_protocol boolean default false,
    organization_name varchar(255) not null,
    organization_fullname varchar(255) not null,
    organization_address varchar(255) not null,
    organization_phone varchar(100),
    organization_fax varchar(100),
    organization_email varchar(100)
);