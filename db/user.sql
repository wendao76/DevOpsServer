CREATE TABLE `users` (
   `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
   `username` VARCHAR(60) NOT NULL COMMENT '用户名(全局唯一)',
   `phone` VARCHAR(20) NULL DEFAULT NULL COMMENT '手机号',
   `email` VARCHAR(60) NULL DEFAULT NULL COMMENT '邮箱',
   `password` VARCHAR(64) NOT NULL COMMENT '密码',
   `nickname` VARCHAR(100) NULL DEFAULT NULL COMMENT '昵称',
   `avatar` VARCHAR(255) NULL DEFAULT NULL COMMENT '头像',
   `created_at` TIMESTAMP NULL DEFAULT NULL COMMENT '创建时间(时间戳)',
   `updated_at` TIMESTAMP NULL DEFAULT NULL COMMENT '更新时间(时间戳)',
   `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间(时间戳)',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
