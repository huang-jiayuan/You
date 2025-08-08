package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// RoomMic 麦位管理表 结构体
type RoomMic struct {
	global.GVA_MODEL
	RoomId      int64     `json:"roomId" form:"roomId" gorm:"comment:关联房间ID;column:room_id;not null;" binding:"required"`
	MicPosition int8      `json:"micPosition" form:"micPosition" gorm:"comment:麦位序号(1-8);column:mic_position;not null;" binding:"required"`
	Status      int8      `json:"status" form:"status" gorm:"comment:麦位状态(0-空闲,1-占用,2-禁用);column:status;not null;default:0;"`
	UserId      uint64    `json:"userId" form:"userId" gorm:"comment:占用麦位的用户ID;column:user_id;default:NULL;"`
	OccupyTime  time.Time `json:"occupyTime" form:"occupyTime" gorm:"comment:占用时间;column:occupy_time;default:NULL;"`
	IsLocked    int8      `json:"isLocked" form:"isLocked" gorm:"comment:是否锁麦(0-未锁,1-已锁);column:is_locked;not null;default:0;"`
	CreatedBy   uint      `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint      `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint      `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RoomMic自定义表名 room_mic
func (RoomMic) TableName() string {
	return "room_mic"
}