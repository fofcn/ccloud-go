create table if not exists `userinfo`(
    `id` integer primary key autoincrement not null,
    `username` varchar(50) not null,
    `password` varchar(128) not null,
    `create_time` datetime not null
    CONSTRAINT "uni_idx_username" UNIQUE ("username")
);

create table if not exists `media_file`(
    `id` integer primary key autoincrement not null,
    `file_name` varchar(128) not null, 
    `store_path` varchar(128) not null,
    `file_create_time` datetime not null,
    `media_type` tinyint not null,
    `create_time` datetime not null
);