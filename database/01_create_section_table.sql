create table home_content.section
(
    id          bigint auto_increment
        primary key,
    title       varchar(255) null,
    description varchar(255) null,
    active      boolean default false,
    deleted     boolean default false
);