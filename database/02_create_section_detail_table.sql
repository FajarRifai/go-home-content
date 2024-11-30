create table home_content.section_detail
(
    id         bigint auto_increment
        primary key,
    code       varchar(100) null,
    rank       int null,
    section_id bigint null,
    constraint section_detail_ibfk_1
        foreign key (section_id) references home_content.section (id)
);

create index section_id
    on home_content.section_detail (section_id);

