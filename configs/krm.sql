-- Date: 2020-07-02

CREATE
DATABASE IF NOT EXISTS krm;

CREATE TABLE `uc_user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `user_id`    varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
    `username`   varchar(253) NOT NULL DEFAULT '' COMMENT '用户名称',
    `status`     varchar(64)  NOT NULL DEFAULT '' COMMENT '用户状态：registered,active,disabled,blacklisted,locked,deleted',
    `nickname`   varchar(253) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `password`   varchar(64)  NOT NULL DEFAULT '' COMMENT '用户加密后的密码',
    `email`      varchar(253) NOT NULL DEFAULT '' COMMENT '用户电子邮箱',
    `phone`      varchar(16)  NOT NULL DEFAULT '' COMMENT '用户手机号',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    UNIQUE KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';


CREATE TABLE `uc_secret`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `user_id`     varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
    `name`        varchar(253) NOT NULL DEFAULT '' COMMENT '密钥名称',
    `secret_id`   varchar(36)  NOT NULL DEFAULT '' COMMENT '密钥 ID',
    `secret_key`  varchar(36)  NOT NULL DEFAULT '' COMMENT '密钥 Key',
    `status`      tinyint unsigned NOT NULL DEFAULT 1 COMMENT '密钥状态，0-禁用；1-启用',
    `expires`     bigint(64) NOT NULL DEFAULT 0 COMMENT '0 永不过期',
    `description` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥描述',
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_secret_id` (`secret_id`),
    KEY           `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='密钥表';
