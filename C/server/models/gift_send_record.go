package models

import (
	"time"
)

type GiftSendRecord struct {
	Id            int64     `gorm:"column:record_id;type:bigint;comment:记录唯一ID;primaryKey;not null;" json:"id"`                                 // 记录唯一ID
	SendUserId    uint64    `gorm:"column:send_user_id;type:bigint UNSIGNED;comment:赠送者用户ID，关联 user.id;not null;" json:"send_user_id"`       // 赠送者用户ID，关联 user.id
	ReceiveUserId uint64    `gorm:"column:receive_user_id;type:bigint UNSIGNED;comment:接收者用户ID，关联 user.id;not null;" json:"receive_user_id"` // 接收者用户ID，关联 user.id
	RoomId        int64     `gorm:"column:room_id;type:bigint;comment:房间ID(NULL-私聊);default:NULL;" json:"room_id"`                              // 房间ID(NULL-私聊)
	GiftId        int64     `gorm:"column:gift_id;type:bigint;comment:礼物ID;not null;" json:"gift_id"`                                             // 礼物ID
	SendCount     int32     `gorm:"column:send_count;type:int;comment:赠送数量;not null;default:1;" json:"send_count"`                              // 赠送数量
	TotalPrice    float64   `gorm:"column:total_price;type:decimal(10, 2);comment:总消耗(虚拟币);not null;" json:"total_price"`                     // 总消耗(虚拟币)
	TotalDiamond  int32     `gorm:"column:total_diamond;type:int;comment:总消耗(钻石);default:0;" json:"total_diamond"`                             // 总消耗(钻石)
	SendType      string    `gorm:"column:send_type;type:varchar(1);comment:赠送方式(1-背包,2-直接购买);not null;" json:"send_type"`                // 赠送方式(1-背包,2-直接购买)
	Message       string    `gorm:"column:message;type:varchar(255);comment:赠送附言;default:NULL;" json:"message"`                                 // 赠送附言
	SendTime      time.Time `gorm:"column:send_time;type:datetime(2);comment:赠送时间;not null;default:CURRENT_TIMESTAMP(2);" json:"send_time"`     // 赠送时间
	Status        string    `gorm:"column:status;type:varchar(1);comment:状态(0-失败,1-成功,2-已撤回);not null;default:1;" json:"status"`           // 状态(0-失败,1-成功,2-已撤回)
	ClientIp      string    `gorm:"column:client_ip;type:varchar(50);comment:赠送者IP;not null;" json:"client_ip"`                                  // 赠送者IP
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);comment:开始时间;not null;" json:"created_at"`                                // 开始时间
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`                                // 修改时间
	CreatedBy     int64     `gorm:"column:created_by;type:bigint;comment:创建者;default:NULL;" json:"created_by"`                                   // 创建者
	UpdatedBy     int64     `gorm:"column:updated_by;type:bigint;comment:修改者;default:NULL;" json:"updated_by"`                                   // 修改者
	DeletedBy     int64     `gorm:"column:deleted_by;type:bigint;comment:删除者;default:NULL;" json:"deleted_by"`                                   // 删除者
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime;comment:删除时间;default:NULL;" json:"deleted_at"`                               // 删除时间
}

func (gi *GiftSendRecord) TableName() string {
	return "gift_send_record"
}
