MySQL表设计
```mysql

create database Registration;

use Registration;

-- User表，存储用户基本信息
CREATE TABLE `User` (
                        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
                        `name` varchar(20) NOT NULL COMMENT '姓名',
                        `student_id` varchar(20) NOT NULL COMMENT '学号',
                        `gender` tinyint(1) NOT NULL COMMENT '性别：0表示女，1表示男',
                        `class` varchar(50) NOT NULL COMMENT '班级',
                        `position` varchar(50) DEFAULT NULL COMMENT '学生职务',
                        `mobile` varchar(20) NOT NULL COMMENT '电话号码',
                        `group` varchar(20) NOT NULL COMMENT '报名组别',
                        `introduction` varchar(500) DEFAULT NULL COMMENT '个人简介',
                        `awards` varchar(500) DEFAULT NULL COMMENT '所获奖项',
                        `reason` varchar(500) DEFAULT NULL COMMENT '加入原因',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                        PRIMARY KEY (`id`)
) COMMENT='用户表';


# -- User表，存储用户基本信息
# CREATE TABLE `User` (
#                         `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
#                         `name` varchar(20) NOT NULL COMMENT '姓名',
#                         `gender` tinyint(1) NOT NULL COMMENT '性别：0表示女，1表示男',
#                         `mobile` varchar(20) NOT NULL COMMENT '手机号码',
#                         `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
#                         `oauth_provider` varchar(20) DEFAULT NULL COMMENT 'OAuth提供商',
#                         `oauth_openid` varchar(50) DEFAULT NULL COMMENT 'OAuth OpenID',
#                         PRIMARY KEY (`id`)
# ) COMMENT='用户表';

-- Application表，存储报名信息
CREATE TABLE `Application` (
                               `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `user_id` int(11) NOT NULL COMMENT '用户ID',
                               `create_time` datetime NOT NULL COMMENT '报名时间',
                               `interview_id` int(11) NOT NULL COMMENT '面试场次ID',
                               `state` tinyint(1) NOT NULL COMMENT '状态：0表示已取消，1表示正常状态',
                               `interview_res` tinyint(1) NOT NULL COMMENT '面试结果：0表示评审中/未面试，-1表示未通过，1表示通过',
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