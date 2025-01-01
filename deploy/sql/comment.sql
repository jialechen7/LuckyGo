DROP DATABASE IF EXISTS `go-lottery-comment`;
CREATE DATABASE `go-lottery-comment` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;

USE `go-lottery-comment`;
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL COMMENT '用户id',
    `lottery_id` int NOT NULL COMMENT '抽奖id',
    `prize_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '奖品名称',
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '晒单评论内容',
    `pics` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '晒单评论图片',
    `praise_count` int NOT NULL COMMENT '点赞数量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '评论表' ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `praise`;
CREATE TABLE `praise` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL COMMENT '点赞用户ID',
    `comment_id` int NOT NULL COMMENT '评论ID',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_comment_unique`(`user_id`, `comment_id`) USING BTREE COMMENT '用户和评论联合唯一索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '点赞表' ROW_FORMAT = DYNAMIC;
