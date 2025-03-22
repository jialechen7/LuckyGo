DROP DATABASE IF EXISTS `go-lottery-lottery`;
CREATE DATABASE `go-lottery-lottery` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;

USE `go-lottery-lottery`;
DROP TABLE IF EXISTS `lottery`;
CREATE TABLE `lottery` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL DEFAULT 0 COMMENT '发起抽奖用户ID',
    `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等奖名称',
    `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等经配图',
    `publish_time` datetime COMMENT '发布抽奖时间',
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
    is_won     TINYINT NOT NULL DEFAULT 0 COMMENT '是否中奖',
    prize_id   BIGINT  NOT NULL DEFAULT 0 COMMENT '中奖id',
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT index_lottery_user UNIQUE (lottery_id, user_id)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '参与表' ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `prize`;
CREATE TABLE `prize`  (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `lottery_id` int(0) NOT NULL DEFAULT 0 COMMENT '抽奖ID',
    `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '奖品类型：1奖品 2优惠券 3兑换码 4商城 5微信红包封面 6红包',
    `name` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品名称',
    `level` int(0) NOT NULL DEFAULT 1 COMMENT '几等奖 默认1',
    `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品图',
    `count` int(0) NOT NULL DEFAULT 0 COMMENT '奖品份数',
    `stock` int(0) NOT NULL DEFAULT 0 COMMENT '奖品库存',
    `grant_type` tinyint(1) NOT NULL COMMENT '奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序',
    `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '奖品表' ROW_FORMAT = Dynamic;