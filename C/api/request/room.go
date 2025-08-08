package request

// KickUser 踢出用户
type KickUser struct {
	RoomID     uint64 `json:"room_id" binding:"required"` // 房间ID
	UserID     uint64 `json:"user_id" binding:"required"` // 被踢用户ID
	OperatorID uint64 `json:"operator_id" binding:"required"`
	Reason     string `json:"reason"` // 原因
}

// MuteUser 禁言
type MuteUser struct {
	RoomID       uint64 `json:"room_id" binding:"required"` // 房间ID
	UserID       uint64 `json:"user_id" binding:"required"` // 被禁言用户ID
	OperatorID   uint64 `json:"operator_id" binding:"required"`
	DurationType int32  `json:"duration_type" binding:"required"` // 禁言时长1小时，2-24小时，3-永久
	Reason       string `json:"reason"`                           // 禁言原因
}

// UnmuteUser 解除禁言
type UnmuteUser struct {
	RoomID     uint64 `json:"room_id" binding:"required"` // 房间ID
	UserID     uint64 `json:"user_id" binding:"required"` // 被解除禁言用户ID
	OperatorID uint64 `json:"operator_id" binding:"required"`
}
