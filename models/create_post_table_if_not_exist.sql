create table if not exists post
(
    id           bigint(20)                               not null auto_increment,
    post_id      bigint(20)                               not null,
    title        varchar(128) collate utf8mb4_general_ci  not null,
    content      varchar(8192) collate utf8mb4_general_ci not null,
    author_id    bigint(20)                               not null,
    community_id bigint(20)                               not null,
    status       tinyint(4)                               not null default 1,
    create_time  timestamp                                not null default current_timestamp,
    update_time  timestamp                                not null default current_timestamp on update current_timestamp,
    primary key (id),
    unique key idx_post_id (post_id),
    key idx_author_id (author_id),
    key idx_community_id (community_id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci