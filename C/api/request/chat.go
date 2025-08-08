package request

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	Type     string  `json:"type" binding:"required"`     // 消息类型: text, image, file, system, typing
	Content  string  `json:"content" binding:"required"`  // 消息内容
	ToUserID uint64  `json:"to_user_id" binding:"required"` // 接收者用户ID
	RoomID   *uint64 `json:"room_id"`                     // 房间ID(群聊时使用)
}

// SystemMessageRequest 系统消息请求
type SystemMessageRequest struct {
	Type    string      `json:"type" binding:"required"`    // 消息类型
	Message string      `json:"message" binding:"required"` // 消息内容
	UserID  *uint64     `json:"user_id"`                    // 目标用户ID，为空则广播
	Data    interface{} `json:"data"`                       // 额外数据
}

// GetHistoryMessagesRequest 获取历史消息请求
type GetHistoryMessagesRequest struct {
	TargetUserID uint64 `json:"target_user_id" binding:"required"` // 目标用户ID
	Limit        int    `json:"limit"`                             // 限制数量，默认20
	Offset       int    `json:"offset"`                            // 偏移量，默认0
}

// MarkMessageAsReadRequest 标记消息已读请求
type MarkMessageAsReadRequest struct {
	MessageID uint64 `json:"message_id" binding:"required"` // 消息ID
}

// GetOfflineMessagesRequest 获取离线消息请求
type GetOfflineMessagesRequest struct {
	Limit  int `json:"limit"`  // 限制数量，默认50
	Offset int `json:"offset"` // 偏移量，默认0
}