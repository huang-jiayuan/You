package request

// 示例
type Streams struct {
}
type CreateRoom struct {
	RoomName string `json:"room_name" form:"room_name"`
	UserId   int    `json:"user_id" form:"user_id" binding:"required"`
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
	Keyword  string `json:"keyword" form:"keyword"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}
