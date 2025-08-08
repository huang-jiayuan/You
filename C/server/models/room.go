package models

import (
	"server/room/basic/global"
	"time"
)

type Room struct {
	Id             int64     `gorm:"column:id;type:bigint;comment:房间唯一ID;primaryKey;not null;" json:"id"`                             // 房间唯一ID
	RoomName       string    `gorm:"column:room_name;type:varchar(50);comment:房间名称;not null;" json:"room_name"`                       // 房间名称
	UserId         uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:房主用户ID，关联 user.id;not null;" json:"user_id"`          // 房主用户ID，关联 user.id
	RoomType       string    `gorm:"column:room_type;type:varchar(1);comment:房间类型(0-公开房,1-私密房);not null;default:0;" json:"room_type"` // 房间类型(0-公开房,1-私密房)
	Status         string    `gorm:"column:status;type:varchar(1);comment:房间状态(0-正常,1-已解散,2-全局禁言);not null;default:0;" json:"status"` // 房间状态(0-正常,1-已解散,2-全局禁言)
	Announcement   string    `gorm:"column:announcement;type:varchar(200);comment:房间公告;default:NULL;" json:"announcement"`            // 房间公告
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	ClosedAt       time.Time `gorm:"column:closed_at;type:datetime;comment:解散时间;default:NULL;" json:"closed_at"`                                       // 解散时间
	LastActiveTime time.Time `gorm:"column:last_active_time;type:datetime;comment:最后活跃时间;not null;default:CURRENT_TIMESTAMP;" json:"last_active_time"` // 最后活跃时间
	FkMemberRoom   int32     `gorm:"column:fk_member_room;type:int;comment:房间人数;not null;default:0;" json:"fk_member_room"`                            // 房间人数
	Image          string    `gorm:"column:image;type:varchar(255);comment:房间封面图;default:NULL;" json:"image"`                                          // 房间封面图
	TagId          uint64    `gorm:"column:tag_id;type:bigint UNSIGNED;comment:关联的标签ID（一个房间只能选一个）;not null;" json:"tag_id"`                            // 关联的标签ID（一个房间只能选一个）
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy      int64     `gorm:"column:created_by;type:bigint;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy      int64     `gorm:"column:updated_by;type:bigint;comment:修改者;default:NULL;" json:"updated_by"` // 修改者
	DeletedBy      int64     `gorm:"column:deleted_by;type:bigint;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (r *Room) TableName() string {
	return "room"
}

func (r *Room) GetFindRoomById(id int64) error {
	return global.DB.Where("id=?", id).First(&r).Error
}
