package models

import (
	"server/user/basic/global"
	"time"
)

// Message 消息表结构体
type Message struct {
	Id         uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:消息唯一ID;primaryKey;" json:"id"`                                    // 消息唯一ID
	MessageType string   `gorm:"column:message_type;type:varchar(20);comment:消息类型(text,image,file,system,typing,heartbeat);not null;default:'text'" json:"message_type"` // 消息类型
	Content     string   `gorm:"column:content;type:text;comment:消息内容;not null;" json:"content"`                                         // 消息内容
	FromUserId  uint64   `gorm:"column:from_user_id;type:bigint UNSIGNED;comment:发送者用户ID;not null;index;" json:"from_user_id"`         // 发送者用户ID
	ToUserId    uint64   `gorm:"column:to_user_id;type:bigint UNSIGNED;comment:接收者用户ID;not null;index;" json:"to_user_id"`             // 接收者用户ID
	RoomId      *uint64  `gorm:"column:room_id;type:bigint UNSIGNED;comment:房间ID(群聊时使用);index;" json:"room_id"`                       // 房间ID(群聊时使用)
	Status      string   `gorm:"column:status;type:varchar(20);comment:消息状态(sent,delivered,read);not null;default:'sent';index;" json:"status"` // 消息状态
	IsOffline   bool     `gorm:"column:is_offline;type:tinyint(1);comment:是否为离线消息;not null;default:0;" json:"is_offline"`              // 是否为离线消息
	ReadAt      *time.Time `gorm:"column:read_at;type:datetime(3);comment:消息已读时间;" json:"read_at"`                                      // 消息已读时间
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime(3);" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime(3);" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:datetime(3);default:null;" json:"deleted_at"`
	CreatedBy   uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;" json:"created_by"` // 创建者
	UpdatedBy   uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;" json:"updated_by"` // 更新者
	DeletedBy   uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;" json:"deleted_by"` // 删除者
}

// TableName 自定义表名
func (m *Message) TableName() string {
	return "messages"
}

// CreateMessage 创建消息
func (m *Message) CreateMessage(fromUserId, toUserId uint64, messageType, content string, roomId *uint64) (*Message, error) {
	message := &Message{
		MessageType: messageType,
		Content:     content,
		FromUserId:  fromUserId,
		ToUserId:    toUserId,
		RoomId:      roomId,
		Status:      "sent",
		IsOffline:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := global.DB.Create(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

// FindMessageById 根据ID查找消息
func (m *Message) FindMessageById(id uint64) (*Message, error) {
	err := global.DB.Where("id = ?", id).First(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// MarkAsRead 标记消息为已读
func (m *Message) MarkAsRead(messageId uint64) error {
	now := time.Now()
	return global.DB.Model(m).Where("id = ?", messageId).Updates(map[string]interface{}{
		"status":  "read",
		"read_at": &now,
	}).Error
}

// MarkAsDelivered 标记消息为已送达
func (m *Message) MarkAsDelivered(messageId uint64) error {
	return global.DB.Model(m).Where("id = ?", messageId).Update("status", "delivered").Error
}

// GetHistoryMessages 获取历史消息
func (m *Message) GetHistoryMessages(userId, targetUserId uint64, limit, offset int) ([]*Message, error) {
	var messages []*Message
	err := global.DB.Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", 
		userId, targetUserId, targetUserId, userId).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, err
}

// GetOfflineMessages 获取离线消息
func (m *Message) GetOfflineMessages(userId uint64) ([]*Message, error) {
	var messages []*Message
	err := global.DB.Where("to_user_id = ? AND is_offline = 1 AND status != 'read'", userId).
		Where("deleted_at IS NULL").
		Order("created_at ASC").
		Find(&messages).Error
	return messages, err
}

// GetUnreadCount 获取未读消息数量
func (m *Message) GetUnreadCount(userId uint64) (int64, error) {
	var count int64
	err := global.DB.Model(m).Where("to_user_id = ? AND status != 'read'", userId).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count, err
}

// IsRead 检查消息是否已读
func (m *Message) IsRead() bool {
	return m.Status == "read"
}

// IsTextMessage 检查是否为文本消息
func (m *Message) IsTextMessage() bool {
	return m.MessageType == "text"
}

// IsSystemMessage 检查是否为系统消息
func (m *Message) IsSystemMessage() bool {
	return m.MessageType == "system"
}

// IsGroupMessage 检查是否为群聊消息
func (m *Message) IsGroupMessage() bool {
	return m.RoomId != nil && *m.RoomId > 0
}