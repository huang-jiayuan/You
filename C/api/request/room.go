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

type ApplyMic struct {
	RoomId int64 `json:"room_id" form:"room_id" binding:"required"`
}

type LeaveMic struct {
	RoomId int64 `json:"room_id" form:"room_id" binding:"required"`
}
type HandleMicApplication struct {
	ApplicationId int64  `json:"application_id" binding:"required"`
	HandlerId     uint64 `json:"handler_id" binding:"required"`
	Action        int32  `json:"action" binding:"required,oneof=1 2"`
	Reason        string `json:"reason"`
}
type KickFromMic struct {
	RoomId       int64  `json:"room_id" form:"room_id" binding:"required"`
	TargetUserId uint64 `json:"target_user_id" form:"target_user_id" binding:"required"`
	Reason       string `json:"reason" form:"reason"`
}
type MuteMicUser struct {
	RoomId       int64  `json:"room_id" form:"room_id" binding:"required"`
	TargetUserId uint64 `json:"target_user_id" form:"target_user_id" binding:"required"`
	Action       int32  `json:"action" form:"action" binding:"required,oneof=1 2"`
	Duration     int32  `json:"duration" form:"duration"`
	Reason       string `json:"reason" form:"reason"`
}

type SendGifts struct {
	SendUserId    int64  `json:"send_user_id" form:"send_user_id"`
	ReceiveUserId int64  `json:"receive_user_id" form:"receive_user_id"`
	RoomId        int64  `json:"room_id" form:"room_id"`
	GiftId        int64  `json:"gift_id" form:"gift_id"`
	SendCount     int64  `json:"send_count" form:"send_count"`
	SendType      string `json:"send_type" form:"send_type"`
	Message       string `json:"message" form:"message"`
	Status        string `json:"status" form:"status"`
	ClientIp      string `json:"client_ip" form:"client_ip"`
	SendTime      string `json:"send_time" form:"send_time"`
}

type SetAdmin struct {
	Id         int64  `json:"id" form:"id"`
	UserId     int64  `json:"user_id" form:"user_id"`
	Status     string `json:"status" form:"status"`
	Reason     string `json:"reason" form:"reason"`
	MuteDay    string `json:"mute_day" form:"mute_day"`
	MuteResult string `json:"mute_result" form:"mute_result"`
	MuteType   string `json:"mute_type" form:"mute_type"`
}
