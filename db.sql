
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

