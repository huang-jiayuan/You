-- 创建麦位管理表
CREATE TABLE IF NOT EXISTS `room_mic` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一ID',
  `room_id` bigint NOT NULL COMMENT '关联房间ID',
  `mic_position` tinyint NOT NULL COMMENT '麦位序号(1-8)',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '麦位状态(0-空闲,1-占用,2-禁用)',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '占用麦位的用户ID，关联 user.id',
  `occupy_time` datetime DEFAULT NULL COMMENT '占用时间',
  `is_locked` tinyint NOT NULL DEFAULT '0' COMMENT '是否锁麦(0-未锁,1-已锁)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_room_mic_position` (`room_id`, `mic_position`),
  KEY `idx_room_mic_room_id` (`room_id`),
  KEY `idx_room_mic_user_id` (`user_id`),
  KEY `idx_room_mic_status` (`status`),
  KEY `idx_room_mic_room_status` (`room_id`, `status`),
  KEY `idx_room_mic_position` (`mic_position`),
  CONSTRAINT `fk_room_mic_room` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_room_mic_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `chk_room_mic_position` CHECK (`mic_position` BETWEEN 1 AND 8),
  CONSTRAINT `chk_room_mic_status` CHECK (`status` IN (0, 1, 2)),
  CONSTRAINT `chk_room_mic_locked` CHECK (`is_locked` IN (0, 1))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='麦位管理表';