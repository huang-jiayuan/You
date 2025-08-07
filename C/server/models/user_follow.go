package models

type UserFollow struct {
	Id         int32 `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	FollowerId int32 `gorm:"column:follower_id;type:int;comment:关注者;not null;" json:"follower_id"`   // 关注者
	FollowedId int32 `gorm:"column:followed_id;type:int;comment:被关注者;not null;" json:"followed_id"` // 被关注者
}

func (uf *UserFollow) TableName() string {
	return "user_follow"
}
