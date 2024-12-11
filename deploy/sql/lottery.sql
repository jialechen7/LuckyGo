DROP DATABASE IF EXISTS `go-lottery-lottery`;
CREATE DATABASE `go-lottery-lottery` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;

USE `go-lottery-lottery`;
DROP TABLE IF EXISTS `lottery`;
CREATE TABLE `lottery` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL DEFAULT 0 COMMENT '发起抽奖用户ID',
    `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等奖名称',
    `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等经配图',
    `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布抽奖时间',
    `join_number` int NOT NULL DEFAULT 0 COMMENT '自动开奖人数',
    `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '抽奖说明',
    `award_deadline` datetime NOT NULL COMMENT '领奖截止时间',
    `is_selected` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否精选: 0否 1是',
    `announce_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '开奖设置：1按时间开奖 2按人数开奖 3即抽即中',
    `announce_time` datetime NOT NULL  COMMENT '开奖时间',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_announced` tinyint(1) NULL DEFAULT 0 COMMENT '是否开奖：0未开奖 1已经开奖',
    `sponsor_id` int NOT NULL DEFAULT 0 COMMENT '发起抽奖赞助商ID',
    `is_clocked` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启打卡任务：0未开启 1已开启',
    `clock_task_id` int NOT NULL DEFAULT 0 COMMENT '打卡任务任务ID',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '抽奖表' ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `lottery_participation`;
CREATE TABLE `lottery_participation` (
    id BIGINT AUTO_INCREMENT COMMENT '主键' PRIMARY KEY,
    lottery_id INT     NOT NULL COMMENT '参与的抽奖的id',
    user_id    INT     NOT NULL COMMENT '用户id',
    is_won     TINYINT NOT NULL COMMENT '是否中奖',
    prize_id   BIGINT  NOT NULL COMMENT '中奖id',
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT index_lottery_user UNIQUE (lottery_id, user_id)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '参与表' ROW_FORMAT = DYNAMIC;

