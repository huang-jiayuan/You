package models

import "server/user/basic/global"

type UserFollow struct {
	Id         int32 `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	FollowerId int64 `gorm:"column:follower_id;type:int;comment:关注者;not null;" json:"follower_id"`   // 关注者
	FollowedId int64 `gorm:"column:followed_id;type:int;comment:被关注者;not null;" json:"followed_id"` // 被关注者
}

func (uf *UserFollow) TableName() string {
	return "user_follow"
}

func (uf *UserFollow) Follow(userId, followedId int64) error {
	follow := &UserFollow{
		FollowerId: userId,
		FollowedId: followedId,
	}
	return global.DB.Create(follow).Error
}
func (uf *UserFollow) Unfollow(userId, followedId int64) error {
	return global.DB.Where("follower_id=? AND followed_id=?", userId, followedId).Delete(&UserFollow{}).Error
}
func (uf *UserFollow) FollowList(userId int64) (userFollow []*UserFollow, err error) {
	err = global.DB.Where("follower_id=?", userId).Find(&userFollow).Error
	if err != nil {
		return nil, err
	}
	return userFollow, nil
}
