package models

import (
	"server/user/basic/global"
	"time"
)

// ChatSession 聊天会话表结构体
type ChatSession struct {
	Id                 uint64     `gorm:"column:id;type:bigint UNSIGNED;comment:会话唯一ID;primaryKey;" json:"id"`                                      // 会话唯一ID
	User1Id            uint64     `gorm:"column:user1_id;type:bigint UNSIGNED;comment:用户1的ID;not null;uniqueIndex:idx_users;" json:"user1_id"`      // 用户1的ID
	User2Id            uint64     `gorm:"column:user2_id;type:bigint UNSIGNED;comment:用户2的ID;not null;uniqueIndex:idx_users;" json:"user2_id"`      // 用户2的ID
	LastMessageId      *uint64    `gorm:"column:last_message_id;type:bigint UNSIGNED;comment:最后一条消息ID;" json:"last_message_id"`                    // 最后一条消息ID
	LastMessageTime    *time.Time `gorm:"column:last_message_time;type:datetime(3);comment:最后消息时间;index;" json:"last_message_time"`                // 最后消息时间
	LastMessageContent string     `gorm:"column:last_message_content;type:varchar(200);comment:最后消息内容预览;" json:"last_message_content"`             // 最后消息内容预览
	User1UnreadCount   int        `gorm:"column:user1_unread_count;type:int;comment:用户1未读消息数;not null;default:0;" json:"user1_unread_count"`      // 用户1未读消息数
	User2UnreadCount   int        `gorm:"column:user2_unread_count;type:int;comment:用户2未读消息数;not null;default:0;" json:"user2_unread_count"`      // 用户2未读消息数
	User1LastReadTime  *time.Time `gorm:"column:user1_last_read_time;type:datetime(3);comment:用户1最后阅读时间;" json:"user1_last_read_time"`            // 用户1最后阅读时间
	User2LastReadTime  *time.Time `gorm:"column:user2_last_read_time;type:datetime(3);comment:用户2最后阅读时间;" json:"user2_last_read_time"`            // 用户2最后阅读时间
	IsActive           bool       `gorm:"column:is_active;type:tinyint(1);comment:会话是否活跃;not null;default:1;" json:"is_active"`                    // 会话是否活跃
	SessionType        string     `gorm:"column:session_type;type:varchar(20);comment:会话类型(private,group);not null;default:'private';" json:"session_type"` // 会话类型
	CreatedAt          time.Time  `gorm:"column:created_at;type:datetime(3);" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at;type:datetime(3);" json:"updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at;type:datetime(3);default:null;" json:"deleted_at"`
	CreatedBy          uint64     `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;" json:"created_by"` // 创建者
	UpdatedBy          uint64     `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;" json:"updated_by"` // 更新者
	DeletedBy          uint64     `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;" json:"deleted_by"` // 删除者
}

// TableName 自定义表名
func (c *ChatSession) TableName() string {
	return "chat_sessions"
}

// CreateChatSession 创建聊天会话
func (c *ChatSession) CreateChatSession(user1Id, user2Id uint64, sessionType string) (*ChatSession, error) {
	// 确保user1Id < user2Id，保持一致性
	if user1Id > user2Id {
		user1Id, user2Id = user2Id, user1Id
	}
	
	session := &ChatSession{
		User1Id:            user1Id,
		User2Id:            user2Id,
		User1UnreadCount:   0,
		User2UnreadCount:   0,
		IsActive:           true,
		SessionType:        sessionType,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	
	err := global.DB.Create(session).Error
	if err != nil {
		return nil, err
	}
	return session, nil
}

// FindOrCreateSession 查找或创建会话
func (c *ChatSession) FindOrCreateSession(user1Id, user2Id uint64) (*ChatSession, error) {
	// 确保user1Id < user2Id，保持一致性
	if user1Id > user2Id {
		user1Id, user2Id = user2Id, user1Id
	}
	
	// 先尝试查找现有会话
	err := global.DB.Where("user1_id = ? AND user2_id = ?", user1Id, user2Id).
		Where("deleted_at IS NULL").
		First(c).Error
	
	if err != nil {
		// 会话不存在，创建新会话
		return c.CreateChatSession(user1Id, user2Id, "private")
	}
	
	return c, nil
}

// GetUserSessions 获取用户的会话列表
func (c *ChatSession) GetUserSessions(userId uint64, limit, offset int) ([]*ChatSession, error) {
	var sessions []*ChatSession
	err := global.DB.Where("(user1_id = ? OR user2_id = ?) AND is_active = 1", userId, userId).
		Where("deleted_at IS NULL").
		Order("last_message_time DESC").
		Limit(limit).
		Offset(offset).
		Find(&sessions).Error
	return sessions, err
}

// UpdateLastMessage 更新最后一条消息信息
func (c *ChatSession) UpdateLastMessage(sessionId, messageId uint64, content string, senderId uint64) error {
	now := time.Now()
	updates := map[string]interface{}{
		"last_message_id":      messageId,
		"last_message_time":    &now,
		"last_message_content": content,
		"updated_at":           now,
	}
	
	// 如果消息内容过长，截断
	if len(content) > 200 {
		updates["last_message_content"] = content[:197] + "..."
	}
	
	// 增加接收者的未读消息数
	var session ChatSession
	err := global.DB.Where("id = ?", sessionId).First(&session).Error
	if err != nil {
		return err
	}
	
	if session.User1Id == senderId {
		// 发送者是user1，增加user2的未读数
		updates["user2_unread_count"] = session.User2UnreadCount + 1
	} else {
		// 发送者是user2，增加user1的未读数
		updates["user1_unread_count"] = session.User1UnreadCount + 1
	}
	
	return global.DB.Model(c).Where("id = ?", sessionId).Updates(updates).Error
}

// ClearUnreadCount 清空指定用户的未读消息数
func (c *ChatSession) ClearUnreadCount(sessionId, userId uint64) error {
	now := time.Now()
	
	var session ChatSession
	err := global.DB.Where("id = ?", sessionId).First(&session).Error
	if err != nil {
		return err
	}
	
	updates := map[string]interface{}{
		"updated_at": now,
	}
	
	if session.User1Id == userId {
		updates["user1_unread_count"] = 0
		updates["user1_last_read_time"] = &now
	} else if session.User2Id == userId {
		updates["user2_unread_count"] = 0
		updates["user2_last_read_time"] = &now
	}
	
	return global.DB.Model(c).Where("id = ?", sessionId).Updates(updates).Error
}

// GetUnreadCount 获取指定用户的未读消息数
func (c *ChatSession) GetUnreadCount(sessionId, userId uint64) (int, error) {
	var session ChatSession
	err := global.DB.Where("id = ?", sessionId).First(&session).Error
	if err != nil {
		return 0, err
	}
	
	if session.User1Id == userId {
		return session.User1UnreadCount, nil
	} else if session.User2Id == userId {
		return session.User2UnreadCount, nil
	}
	
	return 0, nil
}

// GetTotalUnreadCount 获取用户所有会话的未读消息总数
func (c *ChatSession) GetTotalUnreadCount(userId uint64) (int, error) {
	var totalCount int
	
	// 查询用户作为user1的会话
	var count1 int
	err := global.DB.Model(c).Where("user1_id = ? AND is_active = 1", userId).
		Where("deleted_at IS NULL").
		Select("COALESCE(SUM(user1_unread_count), 0)").
		Scan(&count1).Error
	if err != nil {
		return 0, err
	}
	
	// 查询用户作为user2的会话
	var count2 int
	err = global.DB.Model(c).Where("user2_id = ? AND is_active = 1", userId).
		Where("deleted_at IS NULL").
		Select("COALESCE(SUM(user2_unread_count), 0)").
		Scan(&count2).Error
	if err != nil {
		return 0, err
	}
	
	totalCount = count1 + count2
	return totalCount, nil
}

// GetOtherUserId 获取对方用户ID
func (c *ChatSession) GetOtherUserId(currentUserId uint64) uint64 {
	if c.User1Id == currentUserId {
		return c.User2Id
	}
	return c.User1Id
}

// IsUserInSession 检查用户是否在此会话中
func (c *ChatSession) IsUserInSession(userId uint64) bool {
	return c.User1Id == userId || c.User2Id == userId
}

// DeactivateSession 停用会话
func (c *ChatSession) DeactivateSession(sessionId uint64) error {
	return global.DB.Model(c).Where("id = ?", sessionId).Update("is_active", false).Error
}

// ActivateSession 激活会话
func (c *ChatSession) ActivateSession(sessionId uint64) error {
	return global.DB.Model(c).Where("id = ?", sessionId).Update("is_active", true).Error
}

// IsPrivateSession 检查是否为私聊会话
func (c *ChatSession) IsPrivateSession() bool {
	return c.SessionType == "private"
}

// IsGroupSession 检查是否为群聊会话
func (c *ChatSession) IsGroupSession() bool {
	return c.SessionType == "group"
}

// HasUnreadMessages 检查指定用户是否有未读消息
func (c *ChatSession) HasUnreadMessages(userId uint64) bool {
	if c.User1Id == userId {
		return c.User1UnreadCount > 0
	} else if c.User2Id == userId {
		return c.User2UnreadCount > 0
	}
	return false
}