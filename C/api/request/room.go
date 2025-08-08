package request

// 示例
type Streams struct {
}

type SendGifts struct {
	SendUserId    int64  `json:"send_user_id" gorm:"send_user_id" binding:"required"`
	ReceiveUserId int64  `json:"receive_user_id" gorm:"receive_user_id" binding:"required"`
	RoomId        int64  `json:"room_id" gorm:"room_id" binding:"required"`
	GiftId        int64  `json:"gift_id" gorm:"gift_id" binding:"required"`
	SendCount     int64  `json:"send_count" gorm:"send_count" binding:"required"`
	SendType      string `json:"send_type" gorm:"send_type" binding:"required"`
	Message       string `json:"message" gorm:"message" binding:"required"`
	Status        string `json:"status" gorm:"status" binding:"required"`
	ClientIp      string `json:"client_ip" gorm:"client_ip" binding:"required"`
	SendTime      string `json:"send_time" gorm:"send_time" binding:"required"`
}
