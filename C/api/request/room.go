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
type CreateRoom struct {
	RoomName string `json:"room_name" form:"room_name"`
	TagId    int    `json:"tag_id" form:"tag_id" binding:"required"`
	Content  string `json:"content" form:"content"`
	IdCard   string `json:"id_card" form:"id_card" binding:"required"`
	RealName string `json:"real_name" form:"real_name" binding:"required"`
}
type UpdateRoom struct {
	Id           int    `json:"id" form:"id" binding:"required"`
	Announcement string `json:"announcement" form:"announcement" binding:"required"`
}
type GetRecommendRooms struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}
type GetRoomsByCategory struct {
	TagId    int `json:"tag_id" form:"tag_id" binding:"required"`
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}
type SearchRooms struct {
	Keyword  string `json:"keyword" form:"keyword" binding:"required,min=1,max=50"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type JoinRoom struct {
	RoomId int64 `json:"room_id" form:"room_id" `
}

type CloseRoom struct {
	RoomId int64 `json:"room_id" form:"room_id" binding:"required"`
}

// 麦位管理相关请求结构
type ApplyMic struct {
	RoomId int64 `json:"room_id" form:"room_id" binding:"required"`
}

type LeaveMic struct {
	RoomId int64 `json:"room_id" form:"room_id" binding:"required"`
}
