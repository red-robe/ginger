
-- 创建用户
create user 'ginger_user'@'localhost' identified by '123456';
create database ginger_db;
grant all on ginger_db.* to 'ginger_user'@'localhost';

--
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(12) NOT NULL,
  `age` tinyint(3) unsigned zerofill DEFAULT '000',
  `gender` tinyint(1) DEFAULT '0',
  `avatar` varchar(255) DEFAULT '',
  `email` varchar(32) DEFAULT '',
  `phone` char(11) DEFAULT '',
  `password` char(40) NOT NULL,
  `salt` char(4) DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `user_oauth2`;
CREATE TABLE `user_oauth2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platform` tinyint(4) NOT NULL COMMENT '平台账号类型',
  `access_token` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '获取三方平台用户信息的accessToken',
  `open_id` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '开发者可通过OpenID来获取用户基本信息',
  `union_id` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '如果开发者拥有多个移动应用、网站应用和公众帐号，可通过获取用户基本信息中的unionid来区分用户的唯一性，因为只要是同一个微信开放平台帐号下的移动应用、网站应用和公众帐号，用户的unionid是唯一的。',
  `nick_name` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '三方平台账号的用户昵称',
  `gender` tinyint(4) DEFAULT NULL COMMENT '用户性别：1男2女',
  `avatar_url` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户头像',
  `create_at` datetime DEFAULT NULL COMMENT '创建或绑定时间',
  `update_at` datetime DEFAULT NULL COMMENT '最近登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user_oauth2_binding`;
CREATE TABLE `user_oauth2_binding` (
  `user_id` int(11) NOT NULL COMMENT '用户表id',
  `oauth_user_id` int(11) NOT NULL COMMENT '三方用户id绑定',
  UNIQUE KEY `user_oauth2_binding_oauth_user_id_uindex` (`oauth_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;