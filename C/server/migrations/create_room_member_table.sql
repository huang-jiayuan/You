-- 创建房间成员表
CREATE TABLE IF NOT EXISTS `room_member` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一ID',
  `room_id` bigint NOT NULL COMMENT '关联房间ID',
  `user_id` bigint unsigned NOT NULL COMMENT '成员用户ID，关联 user.id',
  `role` tinyint NOT NULL DEFAULT '0' COMMENT '成员角色(0-普通,1-管理员,2-房主)',
  `join_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `leave_time` datetime DEFAULT NULL COMMENT '离开时间',
  `is_muted` tinyint NOT NULL DEFAULT '0' COMMENT '是否禁言(0-正常,1-禁言)',
  `mute_end_time` datetime DEFAULT NULL COMMENT '禁言结束时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_room_member` (`room_id`, `user_id`),
  KEY `idx_room_member_room_id` (`room_id`),
  KEY `idx_room_member_user_id` (`user_id`),
  KEY `idx_room_member_role` (`role`),
  KEY `idx_room_member_muted` (`is_muted`),
  KEY `idx_room_member_room_role` (`room_id`, `role`),
  CONSTRAINT `fk_room_member_room` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_room_member_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `chk_room_member_role` CHECK (`role` IN (0, 1, 2)),
  CONSTRAINT `chk_room_member_muted` CHECK (`is_muted` IN (0, 1))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='房间成员表';