package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// RoomMember 房间成员表 结构体
type RoomMember struct {
	global.GVA_MODEL
	RoomId      int64     `json:"roomId" form:"roomId" gorm:"comment:关联房间ID;column:room_id;not null;" binding:"required"`
	UserId      uint64    `json:"userId" form:"userId" gorm:"comment:成员用户ID;column:user_id;not null;" binding:"required"`
	Role        int8      `json:"role" form:"role" gorm:"comment:成员角色(0-普通,1-管理员,2-房主);column:role;not null;default:0;"`
	IsMuted     int8      `json:"isMuted" form:"isMuted" gorm:"comment:是否禁言(0-正常,1-禁言);column:is_muted;not null;default:0;"`
	MuteEndTime time.Time `json:"muteEndTime" form:"muteEndTime" gorm:"comment:禁言结束时间;column:mute_end_time;default:NULL;"`
	CreatedBy   uint      `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint      `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint      `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RoomMember自定义表名 room_member
func (RoomMember) TableName() string {
	return "room_member"
}