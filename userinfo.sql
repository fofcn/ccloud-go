create table if not exists `userinfo`(
    `id` integer primary key autoincrement not null,
    `username` varchar(50) not null,
    `password` varchar(128) not null,
    `create_time` datetime not null,
    CONSTRAINT "uni_idx_username" UNIQUE ("username")
);