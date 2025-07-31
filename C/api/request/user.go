package request

// 示例
type Stream struct {
}
type SendSms struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Source string `form:"source" json:"source" binding:"required"`
}
type UserLogin struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
	Code   string `json:"code" form:"code" binding:"required"`
}
