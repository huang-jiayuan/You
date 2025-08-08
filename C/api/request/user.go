package request

// 示例
type Stream struct {
}
type SendSms struct {
	Username string `form:"username" json:"username" binding:"required"`
	Source   string `form:"source" json:"source" binding:"required"`
}
type UserLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
}

type UserPassword struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}
type UpdatePassword struct {
	Password string `json:"password" form:"password" binding:"required"`
}

type ImproveUserMessage struct {
	Nickname    string `json:"nickname" form:"nickname"`
	Gender      string `json:"gender" form:"gender"`
	VoiceTag    string `json:"voice_tag" form:"voice_tag"`
	InterestTag string `json:"interest_tag" form:"interest_tag"`
}

type FollowUser struct {
	FollowedId int64 `json:"followed_id" binding:"required"`
}

type UnFollowUser struct {
	FollowedId int64 `json:"followed_id" binding:"required"`
}
