DROP DATABASE IF EXISTS `go-lottery-usercenter`;
CREATE DATABASE `go-lottery-usercenter` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;

USE `go-lottery-usercenter`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
    `id` bigint(0) NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
    `mobile` char(11) DEFAULT '' COMMENT '手机号',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
    `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
    `sex` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别 0:男 1:女',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `info` varchar(255) NOT NULL DEFAULT '' COMMENT '简介',
    `is_admin` tinyint(1) DEFAULT 0 COMMENT '是否管理员 0:否 1:是',
    `signature` varchar(200) NOT NULL DEFAULT '' COMMENT '个性签名',
    `location_name` varchar(100) NOT NULL DEFAULT '' COMMENT '地址名称',
    `longitude` DOUBLE PRECISION NOT NULL DEFAULT 0 COMMENT '经度',
    `latitude` DOUBLE PRECISION NOT NULL DEFAULT 0 COMMENT '纬度',
    `total_prize` int(0) NOT NULL DEFAULT 0 COMMENT '累计奖品',
    `fans` int(0) NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    `all_lottery` int(0) NOT NULL DEFAULT 0 COMMENT '全部抽奖包含我发起的、我中奖的',
    `initiation_record` int(0) NOT NULL DEFAULT 0 COMMENT '发起抽奖记录',
    `winning_record` int(0) NOT NULL DEFAULT 0 COMMENT '中奖记录',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth`  (
    `id` bigint(0) NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `user_id` bigint(0) NOT NULL DEFAULT 0,
    `auth_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台唯一id',
    `auth_type` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台类型',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_type_key`(`auth_type`, `auth_key`) USING BTREE,
    UNIQUE INDEX `idx_userId_key`(`user_id`, `auth_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户授权表' ROW_FORMAT = Dynamic;
