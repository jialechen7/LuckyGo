DROP DATABASE IF EXISTS `go-lottery-upload`;
CREATE DATABASE `go-lottery-upload` CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;

USE `go-lottery-upload`;
DROP TABLE IF EXISTS `upload_file`;
create table `upload_file` (
    id int NOT NULL primary key auto_increment,
    user_id int NOT NULL comment '上传用户id',
    file_name varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL comment '文件名',
    ext varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL comment '扩展名',
    size int NOT NULL comment '文件大小',
    url varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL comment '下载链接',
    create_time datetime NOT NULL default current_timestamp,
    update_time datetime NOT NULL default current_timestamp ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文件上传表';
