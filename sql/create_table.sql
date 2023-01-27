create table m_users
(
    uuid         uuid                     default gen_random_uuid() not null
        constraint m_users_pk
            primary key,
    email        varchar(255)                                       not null,
    password     varchar(255)                                       not null,
    company_id   integer                                            not null,
    created_time timestamp with time zone default now()             not null,
    updated_time timestamp with time zone default now(),
    deleted_time timestamp with time zone
);

alter table m_users
    owner to shi;

create unique index m_users_uuid_uindex
    on m_users (uuid);

