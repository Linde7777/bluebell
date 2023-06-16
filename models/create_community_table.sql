drop table community;
create table community
(
    id             int(11)                                 not null auto_increment,
    community_id   int(10) unsigned                        not null,
    community_name varchar(128) collate utf8mb4_general_ci not null,
    introduction   varchar(256) collate utf8mb4_general_ci not null,
    create_time    timestamp                               not null default current_timestamp,
    update_time    timestamp                               not null default current_timestamp on update current_timestamp,
    primary key (id),
    unique key idx_community_id (community_id),
    unique key idx_community_name (community_name)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci;

insert into community
values ('1', '1', 'Go', 'Let\'s learn Golang', '2019-11-11 09:35:16', '2019-11-11 09:39:17');


insert into community
values ('2', '2', 'RDR2', 'Red Dead Redemption II is a game developed by Rockstar Games', '2018-10-21 09:35:36',
        '2019-07-15 10:21:19');

insert into community
values ('3', '3', 'GTA5', 'Grand Theft Auto V is a game developed by Rockstar Games', '2010-8-11 17:35:16',
        '2022-05-03 19:06:33');

insert into community
values ('4', '4', 'GTA6', 'Coming soon!', '2022-06-16 15:10:26', '2022-06-16 19:32:11');




