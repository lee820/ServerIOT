drop database if exists iot_service;

CREATE DATABASE iot_service DEFAULT CHARACTER
SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

/*
 * 切换到iot_service数据库
 */
use iot_service;

-- 创建用户表
drop table if exists user;
CREATE TABLE `user` (
	`id`          int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name`        varchar(16) NOT NULL UNIQUE COMMENT '用户名',
	`password`    varchar(255) NOT NULL COMMENT '用户密码',
	`phone`       char(11) NOT NULL UNIQUE COMMENT '用户手机号',
	`dev_count`   int NOT NULL DEFAULT '0' COMMENT '用户当前拥有设备数量',
	`created_on`  int unsigned DEFAULT '0' COMMENT '创建时间',
	`modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`      tinyint unsigned DEFAULT '0' COMMENT '是否删除 0为未删除， 1为已删除'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 创建用户设备关联表
/* 暂不需要
drop table if exists user_device;
CREATE TABLE `user_device` (
	`id`          int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`user_id`     int unsigned NOT NULL COMMENT '用户ID',
	`device_id`   int unsigned NOT NULL COMMENT '设备ID',
	`created_on`  int unsigned DEFAULT '0' COMMENT '创建时间',
	`modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`      tinyint unsigned DEFAULT '0' COMMENT '是否删除 0为未删除， 1为已删除'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户设备关联表';
*/
-- 创建设备表
drop table if exists device;
CREATE TABLE `device` (
	`id`          int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`user_id`     int unsigned NOT NULL COMMENT '所属的用户ID',
	`name`        varchar(16) NOT NULL COMMENT '设备名称',
	`place`       varchar(255) DEFAULT '' COMMENT '设备安放的位置',
	`running`     tinyint unsigned DEFAULT '0' COMMENT '设备运行状态，0表示未运行，1表示运行中',
	`created_on`  int unsigned DEFAULT '0' COMMENT '创建时间',
	`modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`      tinyint unsigned DEFAULT '0' COMMENT '是否删除 0为未删除， 1为已删除'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备表';

-- 创建设备数据表
drop table if exists device_data;
CREATE TABLE `device_data` (
	`id`          int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`device_id`   int unsigned NOT NULL COMMENT '设备ID',
	`device_data` int NOT NULL COMMENT '设备数据',
	`created_on`  int unsigned DEFAULT '0' COMMENT '创建时间',
	`modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`      tinyint unsigned DEFAULT '0' COMMENT '是否删除 0为未删除， 1为已删除'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备数据表';