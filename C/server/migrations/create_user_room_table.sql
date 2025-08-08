-- 创建用户房间关系表
CREATE TABLE IF NOT EXISTS `user_room` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID，关联 user.id',
  `room_id` bigint NOT NULL COMMENT '房间ID，关联 room.id',
  `is_online` tinyint NOT NULL DEFAULT '0' COMMENT '是否在线(0-离线,1-在线)',
  `join_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '进入时间',
  `leave_time` datetime DEFAULT NULL COMMENT '离开时间',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_room_user_id` (`user_id`),
  KEY `idx_user_room_room_id` (`room_id`),
  KEY `idx_user_room_online` (`room_id`, `is_online`),
  UNIQUE KEY `uk_user_room_online` (`user_id`, `room_id`, `is_online`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户房间关系表';