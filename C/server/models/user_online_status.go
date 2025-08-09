package models

import (
	"server/user/basic/global"
	"time"
)

// UserOnlineStatus 用户在线状态表结构体
type UserOnlineStatus struct {
	Id                 uint64     `gorm:"column:id;type:bigint UNSIGNED;comment:状态记录ID;primaryKey;" json:"id"`                                    // 状态记录ID
	UserId             uint64     `gorm:"column:user_id;type:bigint UNSIGNED;comment:用户ID;not null;uniqueIndex;" json:"user_id"`                    // 用户ID
	IsOnline           bool       `gorm:"column:is_online;type:tinyint(1);comment:是否在线;not null;default:0;index;" json:"is_online"`               // 是否在线
	LastActiveTime     time.Time  `gorm:"column:last_active_time;type:datetime(3);comment:最后活跃时间;not null;index;" json:"last_active_time"`      // 最后活跃时间
	WebSocketSessionId *string    `gorm:"column:websocket_session_id;type:varchar(64);comment:WebSocket会话ID;" json:"websocket_session_id"`          // WebSocket会话ID
	DeviceType         string     `gorm:"column:device_type;type:varchar(20);comment:设备类型(web,mobile,desktop);default:'web';" json:"device_type"` // 设备类型
	IpAddress          string     `gorm:"column:ip_address;type:varchar(45);comment:IP地址;" json:"ip_address"`                                       // IP地址
	UserAgent          string     `gorm:"column:user_agent;type:varchar(500);comment:用户代理;" json:"user_agent"`                                    // 用户代理
	CreatedAt          time.Time  `gorm:"column:created_at;type:datetime(3);" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at;type:datetime(3);" json:"updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at;type:datetime(3);default:null;" json:"deleted_at"`
	CreatedBy          uint64     `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;" json:"created_by"` // 创建者
	UpdatedBy          uint64     `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;" json:"updated_by"` // 更新者
	DeletedBy          uint64     `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;" json:"deleted_by"` // 删除者
}

// TableName 自定义表名
func (u *UserOnlineStatus) TableName() string {
	return "user_online_status"
}

// SetUserOnline 设置用户为在线状态
func (u *UserOnlineStatus) SetUserOnline(userId uint64, sessionId, deviceType, ipAddress, userAgent string) error {
	now := time.Now()
	status := &UserOnlineStatus{
		UserId:             userId,
		IsOnline:           true,
		LastActiveTime:     now,
		WebSocketSessionId: &sessionId,
		DeviceType:         deviceType,
		IpAddress:          ipAddress,
		UserAgent:          userAgent,
		UpdatedAt:          now,
	}

	// 使用 ON DUPLICATE KEY UPDATE 或者先查询再更新
	var existingStatus UserOnlineStatus
	err := global.DB.Where("user_id = ?", userId).First(&existingStatus).Error
	if err != nil {
		// 记录不存在，创建新记录
		status.CreatedAt = now
		return global.DB.Create(status).Error
	} else {
		// 记录存在，更新记录
		return global.DB.Model(&existingStatus).Updates(map[string]interface{}{
			"is_online":            true,
			"last_active_time":     now,
			"websocket_session_id": &sessionId,
			"device_type":          deviceType,
			"ip_address":           ipAddress,
			"user_agent":           userAgent,
			"updated_at":           now,
		}).Error
	}
}

// SetUserOffline 设置用户为离线状态
func (u *UserOnlineStatus) SetUserOffline(userId uint64) error {
	now := time.Now()
	return global.DB.Model(u).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"is_online":            false,
		"last_active_time":     now,
		"websocket_session_id": nil,
		"updated_at":           now,
	}).Error
}

// UpdateLastActive 更新最后活跃时间
func (u *UserOnlineStatus) UpdateLastActive(userId uint64) error {
	now := time.Now()
	return global.DB.Model(u).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"last_active_time": now,
		"updated_at":       now,
	}).Error
}

// GetUserOnlineStatus 获取用户在线状态
func (u *UserOnlineStatus) GetUserOnlineStatus(userId uint64) (*UserOnlineStatus, error) {
	err := global.DB.Where("user_id = ?", userId).First(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetOnlineUsers 获取在线用户列表
func (u *UserOnlineStatus) GetOnlineUsers(limit, offset int) ([]*UserOnlineStatus, error) {
	var users []*UserOnlineStatus
	err := global.DB.Where("is_online = 1").
		Where("deleted_at IS NULL").
		Order("last_active_time DESC").
		Limit(limit).
		Offset(offset).
		Find(&users).Error
	return users, err
}

// GetOnlineUserCount 获取在线用户数量
func (u *UserOnlineStatus) GetOnlineUserCount() (int64, error) {
	var count int64
	err := global.DB.Model(u).Where("is_online = 1").
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count, err
}

// IsUserOnline 检查用户是否在线
func (u *UserOnlineStatus) IsUserOnline(userId uint64) (bool, error) {
	var status UserOnlineStatus
	err := global.DB.Where("user_id = ? AND is_online = 1", userId).First(&status).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsActiveWithin 检查用户是否在指定时间内活跃
func (u *UserOnlineStatus) IsActiveWithin(duration time.Duration) bool {
	return time.Since(u.LastActiveTime) <= duration
}

// GetOnlineDuration 获取在线时长
func (u *UserOnlineStatus) GetOnlineDuration() time.Duration {
	if !u.IsOnline {
		return 0
	}
	return time.Since(u.LastActiveTime)
}

// IsWebDevice 检查是否为Web设备
func (u *UserOnlineStatus) IsWebDevice() bool {
	return u.DeviceType == "web"
}

// IsMobileDevice 检查是否为移动设备
func (u *UserOnlineStatus) IsMobileDevice() bool {
	return u.DeviceType == "mobile"
}

// IsDesktopDevice 检查是否为桌面设备
func (u *UserOnlineStatus) IsDesktopDevice() bool {
	return u.DeviceType == "desktop"
}

// BatchSetOffline 批量设置用户离线
func (u *UserOnlineStatus) BatchSetOffline(userIds []uint64) error {
	now := time.Now()
	return global.DB.Model(u).Where("user_id IN ?", userIds).Updates(map[string]interface{}{
		"is_online":            false,
		"last_active_time":     now,
		"websocket_session_id": nil,
		"updated_at":           now,
	}).Error
}
