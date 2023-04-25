# MySQL数据库表设计

## 说明
有些应该使用tinyint(1)的字段，直接使用了int

## User表

存储用户基本信息

| 字段名          | 类型           | 是否为空 | 注释                  |
|--------------|--------------|------|---------------------|
| id           | int(11)      | 否    | 主键                  |
| open_id      | varchar(20)  | 否    | OAuth2标识            |
| name         | varchar(20)  | 否    | 姓名                  |
| student_id   | varchar(20)  | 否    | 学号                  |
| avatar       | varchar(100) | 是    | 头像地址                |
| is_elab_er   | tinyint(2)   | 否    | 是不是科中的同学：0表示不是，1表示是 |
| gender       | tinyint(2)   | 否    | 性别：0表示女，1表示男        |
| class        | varchar(50)  | 否    | 班级                  |
| position     | varchar(50)  | 是    | 学生职务                |
| mobile       | varchar(20)  | 否    | 电话号码                |
| mail         | varchar(20)  | 否    | 邮箱                  |
| group        | varchar(20)  | 否    | 报名组别                |
| introduction | varchar(500) | 是    | 个人简介                |
| awards       | varchar(500) | 是    | 所获奖项                |
| reason       | varchar(500) | 是    | 加入原因                |
| created_at   | datetime     | 否    | 创建时间                |
| updated_at   | datetime     | 否    | 最后修改时间              |

## Application表

存储报名信息

| 字段名           | 类型         | 是否为空 | 注释                            |
|---------------|------------|------|-------------------------------|
| id            | int(11)    | 否    | 主键                            |
| user_id       | int(11)    | 否    | 用户ID                          |
| create_time   | datetime   | 否    | 报名时间                          |
| interview_id  | int(11)    | 否    | 面试场次ID                        |
| state         | tinyint(2) | 否    | 状态：0表示已取消，1表示正常状态             |
| interview_res | tinyint(2) | 否    | 面试结果：0表示评审中/未面试，-1表示未通过，1表示通过 |
| created_at    | datetime   | 否    | 创建时间                          |
| updated_at    | datetime   | 否    | 最后修改时间                        |

## InterviewSession表

存储面试场次信息

| 字段名         | 类型          | 是否为空 | 注释     |
|-------------|-------------|------|--------|
| id          | int(11)     | 否    | 主键     |
| start_time  | datetime    | 否    | 面试开始时间 |
| end_time    | datetime    | 否    | 面试结束时间 |
| location    | varchar(50) | 否    | 面试地点   |
| capacity    | int(11)     | 否    | 可参加人数  |
| applied_num | int(11)     | 否    | 已报名人数  |
| created_at  | datetime    | 否    | 创建时间   |
| updated_at  | datetime    | 否    | 最后修改时间 |

## DDL语句
如果上面的表格与DDL语句不一致，请以DDL语句为准。
```mysql

create database Registration;

use Registration;
drop table if exists User;
-- User表，存储用户基本信息
CREATE TABLE `User` (
                        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
                        `open_id` varchar(20) NOT NULL COMMENT 'OAuth2标识',
                        `name` varchar(20) NOT NULL COMMENT '姓名',
                        `student_id` varchar(20) NOT NULL COMMENT '学号',
                        `avatar` varchar(100) DEFAULT NULL COMMENT '头像地址',
                        `isELABer` tinyint(2) NOT NULL COMMENT '是不是科中的同学：0表示不是，1表示是',
                        `gender` tinyint(2) NOT NULL COMMENT '性别：0表示女，1表示男',
                        `class` varchar(50) NOT NULL COMMENT '班级',
                        `position` varchar(50) DEFAULT NULL COMMENT '学生职务',
                        `mobile` varchar(20) NOT NULL COMMENT '电话号码',
                        `mail` varchar(20) NOT NULL COMMENT '邮箱',
                        `group` varchar(20) NOT NULL COMMENT '报名组别',
                        `introduction` varchar(500) DEFAULT NULL COMMENT '个人简介',
                        `awards` varchar(500) DEFAULT NULL COMMENT '所获奖项',
                        `reason` varchar(500) DEFAULT NULL COMMENT '加入原因',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                        PRIMARY KEY (`id`)
) COMMENT='用户表';


-- Application表，存储报名信息
CREATE TABLE `Application` (
                               `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `user_id` int(11) NOT NULL COMMENT '用户ID',
                               `create_time` datetime NOT NULL COMMENT '报名时间',
                               `interview_id` int(11) NOT NULL COMMENT '面试场次ID',
                               `state` tinyint(2) NOT NULL COMMENT '状态：0表示已取消，1表示正常状态',
                               `interview_res` tinyint(2) NOT NULL COMMENT '面试结果：0表示评审中/未面试，-1表示未通过，1表示通过',
                               `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                               PRIMARY KEY (`id`),
                               KEY `idx_user_id` (`user_id`),
                               KEY `idx_interview_id` (`interview_id`)
) COMMENT='报名表';


-- InterviewSession表，存储面试场次信息
CREATE TABLE `InterviewSession` (
                                    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
                                    `start_time` datetime NOT NULL COMMENT '面试开始时间',
                                    `end_time` datetime NOT NULL COMMENT '面试结束时间',
                                    `location` varchar(50) NOT NULL COMMENT '面试地点',
                                    `capacity` int(11) NOT NULL COMMENT '可参加人数',
                                    `applied_num` int(11) NOT NULL DEFAULT '0' COMMENT '已报名人数',
                                    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                                    PRIMARY KEY (`id`)
) COMMENT='面试场次表';

```