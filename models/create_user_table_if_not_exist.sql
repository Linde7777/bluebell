use bluebell;
CREATE TABLE if not exists user
(
    id          BIGINT(20)  NOT NULL AUTO_INCREMENT,
    user_id     BIGINT(20)  NOT NULL,
    username    VARCHAR(64) NOT NULL COLLATE utf8mb4_general_ci,
    password    varchar(64) NOT NULL COLLATE utf8mb4_general_ci,
    email       VARCHAR(64) COLLATE utf8mb4_general_ci,
    gender      TINYINT(4)  NOT NULL DEFAULT '0',
    create_time TIMESTAMP   NULL     DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP   NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY idx_username (username) USING BTREE,
    UNIQUE KEY idx_user_id (user_id) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
