package models

import "errors"

// 聊天模块相关错误定义
var (
	// Message 相关错误
	ErrEmptyContent        = errors.New("消息内容不能为空")
	ErrInvalidFromUser     = errors.New("发送者用户ID无效")
	ErrInvalidToUser       = errors.New("接收者用户ID无效")
	ErrInvalidMessageType  = errors.New("消息类型无效")
	ErrMessageNotFound     = errors.New("消息不存在")
	ErrMessageAlreadyRead  = errors.New("消息已经被标记为已读")

	// UserOnlineStatus 相关错误
	ErrInvalidUserID       = errors.New("用户ID无效")
	ErrInvalidDeviceType   = errors.New("设备类型无效")
	ErrUserNotFound        = errors.New("用户不存在")
	ErrUserAlreadyOnline   = errors.New("用户已经在线")
	ErrUserAlreadyOffline  = errors.New("用户已经离线")

	// ChatSession 相关错误
	ErrSameUser            = errors.New("不能与自己创建聊天会话")
	ErrInvalidSessionType  = errors.New("会话类型无效")
	ErrSessionNotFound     = errors.New("聊天会话不存在")
	ErrSessionAlreadyExists = errors.New("聊天会话已存在")
	ErrSessionInactive     = errors.New("聊天会话已停用")

	// 通用错误
	ErrDatabaseOperation   = errors.New("数据库操作失败")
	ErrInvalidParameter    = errors.New("参数无效")
	ErrPermissionDenied    = errors.New("权限不足")
	ErrInternalServer      = errors.New("服务器内部错误")
)

// 消息类型常量
const (
	MessageTypeText      = "text"
	MessageTypeImage     = "image"
	MessageTypeFile      = "file"
	MessageTypeSystem    = "system"
	MessageTypeTyping    = "typing"
	MessageTypeHeartbeat = "heartbeat"
)

// 消息状态常量
const (
	MessageStatusSent      = "sent"
	MessageStatusDelivered = "delivered"
	MessageStatusRead      = "read"
)

// 设备类型常量
const (
	DeviceTypeWeb     = "web"
	DeviceTypeMobile  = "mobile"
	DeviceTypeDesktop = "desktop"
)

// 会话类型常量
const (
	SessionTypePrivate = "private"
	SessionTypeGroup   = "group"
)

// ValidateMessageType 验证消息类型是否有效
func ValidateMessageType(messageType string) bool {
	validTypes := []string{
		MessageTypeText,
		MessageTypeImage,
		MessageTypeFile,
		MessageTypeSystem,
		MessageTypeTyping,
		MessageTypeHeartbeat,
	}
	for _, validType := range validTypes {
		if messageType == validType {
			return true
		}
	}
	return false
}

// ValidateMessageStatus 验证消息状态是否有效
func ValidateMessageStatus(status string) bool {
	validStatuses := []string{
		MessageStatusSent,
		MessageStatusDelivered,
		MessageStatusRead,
	}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

// ValidateDeviceType 验证设备类型是否有效
func ValidateDeviceType(deviceType string) bool {
	validTypes := []string{
		DeviceTypeWeb,
		DeviceTypeMobile,
		DeviceTypeDesktop,
	}
	for _, validType := range validTypes {
		if deviceType == validType {
			return true
		}
	}
	return false
}

// ValidateSessionType 验证会话类型是否有效
func ValidateSessionType(sessionType string) bool {
	validTypes := []string{
		SessionTypePrivate,
		SessionTypeGroup,
	}
	for _, validType := range validTypes {
		if sessionType == validType {
			return true
		}
	}
	return false
}