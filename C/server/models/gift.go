package models

import (
	"server/room/basic/global"
	"time"
)

type GiftInfo struct {
	GiftId       int64     `gorm:"column:gift_id;type:bigint;comment:礼物唯一ID;primaryKey;not null;" json:"gift_id"`                   // 礼物唯一ID
	GiftName     string    `gorm:"column:gift_name;type:varchar(100);comment:礼物名称;not null;" json:"gift_name"`                      // 礼物名称
	GiftType     string    `gorm:"column:gift_type;type:varchar(1);comment:礼物类型(1-普通,2-特效,3-稀有,4-自定义);default:0;" json:"gift_type"` // 礼物类型(1-普通,2-特效,3-稀有,4-自定义)
	Price        float64   `gorm:"column:price;type:double;comment:虚拟币价格;default:NULL;" json:"price"`                               // 虚拟币价格
	DiamondPrice int64     `gorm:"column:diamond_price;type:bigint;comment:钻石价格(可为空);default:NULL;" json:"diamond_price"`           // 钻石价格(可为空)
	IconUrl      string    `gorm:"column:icon_url;type:varchar(255);comment:礼物图标URL;not null;" json:"icon_url"`                     // 礼物图标URL
	AnimationUrl string    `gorm:"column:animation_url;type:varchar(255);comment:动画URL(特效礼物);default:NULL;" json:"animation_url"`   // 动画URL(特效礼物)
	Weight       int64     `gorm:"column:weight;type:bigint;comment:排序权重;default:NULL;" json:"weight"`                              // 排序权重
	IsHot        string    `gorm:"column:is_hot;type:varchar(1);comment:是否热门(1-是,0-否);default:0;" json:"is_hot"`                    // 是否热门(1-是,0-否)
	IsLimit      string    `gorm:"column:is_limit;type:varchar(1);comment:是否限时(1-是,0-否);default:0;" json:"is_limit"`                // 是否限时(1-是,0-否)
	Status       string    `gorm:"column:status;type:varchar(1);comment:状态(0-下架,1-上架,2-待审核);default:1;" json:"status"`              // 状态(0-下架,1-上架,2-待审核)
	CreatedBy    uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"`              // 创建者
	UpdatedBy    uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"`              // 更新者
	DeletedBy    uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"`              // 删除者
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`                     // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`                     // 修改时间
}

func (g *GiftInfo) TableName() string {
	return "gift_info"
}

func (g *GiftInfo) GetFindGiftById(id int64) error {
	return global.DB.Where("id=?", id).Find(&g).Error
}
