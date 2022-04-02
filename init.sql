CREATE DATABASE IF NOT EXISTS example_docker_db DEFAULT CHARACTER SET utf8mb4;
USE example_docker_db

CREATE TABLE IF NOT EXISTS `user_tab` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0,
  `user_name` varchar(128) NOT NULL DEFAULT '',
  `create_time` bigint(20) NOT NULL DEFAULT 0,
  `modify_time` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`)
);
