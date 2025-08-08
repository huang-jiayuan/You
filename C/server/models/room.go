package models

import (
	"server/user/basic/global"
	"time"
)

// RoomKick 踢出用户记录表
type RoomKick struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomID     uint64    `gorm:"not null;index" json:"room_id"`
	UserID     uint64    `gorm:"not null;index" json:"user_id"`
	OperatorID uint64    `gorm:"not null" json:"operator_id"`
	Reason     string    `gorm:"type:varchar(500)" json:"reason"`
	KickTime   time.Time `gorm:"not null" json:"kick_time"`
	ExpireTime time.Time `gorm:"not null;index" json:"expire_time"` // 10分钟后可重新进入
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// RoomMute 禁言记录
type RoomMute struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomID       uint64     `gorm:"not null;index" json:"room_id"`
	UserID       uint64     `gorm:"not null;index" json:"user_id"`
	OperatorID   uint64     `gorm:"not null" json:"operator_id"`
	DurationType int32      `gorm:"not null" json:"duration_type"` // 1-1小时，2-24小时，3-永久
	Reason       string     `gorm:"type:varchar(500)" json:"reason"`
	MuteTime     time.Time  `gorm:"not null" json:"mute_time"`
	ExpireTime   *time.Time `gorm:"index" json:"expire_time"` // 永久禁言时为null
	IsActive     bool       `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (RoomKick) TableName() string {
	return "room_kick"
}

func (RoomMute) TableName() string {
	return "room_mute"
}

// IsUserKicked 检查用户是否被踢且在冷却时间内
func IsUserKicked(roomID, userID uint64) bool {
	var count int64
	global.DB.Model(&RoomKick{}).Where("room_id = ? and user_id = ? and expire_time > ?",
		roomID, userID, time.Now()).Count(&count)
	return count > 0
}

// IsUserMuted 判断用户是否被禁言
func IsUserMuted(roomID, userID uint64) bool {
	var count int64
	now := time.Now()

	global.DB.Model(&RoomMute{}).Where("room_id = ? and user_id = ? and is_active = ? and (expire_time IS NULL OR expire_time > ?)",
		roomID, userID, true, now).Count(&count)
	return count > 0
}

// GetMuteDuration 根据类型获取禁言时长
func GetMuteDuration(durationType int32) *time.Time {
	now := time.Now()
	switch durationType {
	case 1: // 1小时
		expireTime := now.Add(1 * time.Hour)
		return &expireTime
	case 2: // 24小时
		expireTime := now.Add(24 * time.Hour)
		return &expireTime
	case 3: // 永久
		return nil
	default:
		expireTime := now.Add(1 * time.Hour)
		return &expireTime
	}
}
