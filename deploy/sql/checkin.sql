DROP DATABASE IF EXISTS `go-lottery-checkin`;
CREATE DATABASE `go-lottery-checkin` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;

USE `go-lottery-checkin`;
DROP TABLE IF EXISTS `checkin_record`;
CREATE TABLE `checkin_record`  (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `continuous_checkin_days` tinyint NOT NULL DEFAULT 0 COMMENT 'Number of consecutive check-in days',
    `state` tinyint NOT NULL DEFAULT 0 COMMENT 'Whether to sign in on the day, 1 means signed, 0 means not signed.',
    `last_checkin_date` datetime NULL DEFAULT NULL COMMENT 'Date of last check-in',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `integral`;
CREATE TABLE `integral`  (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `integral` int NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `integral_record`;
CREATE TABLE `integral_record`  (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `integral` int NOT NULL COMMENT 'points added or subtracted',
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_cs_0900_ai_ci NOT NULL,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `task_progress`;
CREATE TABLE `task_progress`  (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `is_participated_lottery` int NOT NULL DEFAULT 0,
    `is_initiated_lottery` int NOT NULL DEFAULT 0,
    `is_sub_checkin` int NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `task_record`;
CREATE TABLE `task_record`  (
    `id` int NOT NULL AUTO_INCREMENT,
    `type` tinyint NOT NULL,
    `user_id` int NOT NULL,
    `task_id` int NOT NULL,
    `is_finished` tinyint NOT NULL DEFAULT 1 COMMENT '0 means not completed, 1 means completed',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks`  (
    `id` int NOT NULL AUTO_INCREMENT,
    `type` tinyint NOT NULL COMMENT '1 for novice, 2 for daily, 3 for weekly',
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_cs_0900_ai_ci NOT NULL,
    `integral` int NOT NULL COMMENT 'Increased wish value after completion',
    `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_cs_0900_ai_ci NULL DEFAULT NULL,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = DYNAMIC;

# 固定任务列表
INSERT INTO `tasks` (type, content, integral, path) VALUES
(1, '参与一次抽奖', 10, '/pages/index/home'),
(1, '发起一次抽奖', 10, '/pages/index/lottery'),
(1, '订阅签到', 15, '');
