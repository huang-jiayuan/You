package models

import "time"

type AdminUser struct {
	Id            int32     `gorm:"column:id;type:int;comment:管理员ID;primaryKey;not null;" json:"id"`                           // 管理员ID
	UserId        int64     `gorm:"column:user_id;type:bigint;comment:用户的ID;default:NULL;" json:"user_id"`                     // 用户的ID
	AdminPassword string    `gorm:"column:admin_password;type:varchar(32);comment:管理员的密码;not null;" json:"admin_password"`     // 管理员的密码
	AdminName     string    `gorm:"column:admin_name;type:varchar(25);comment:管理员账号;not null;" json:"admin_name"`              // 管理员账号
	Status        string    `gorm:"column:status;type:varchar(2);comment:用户的账号状态(0-封禁,1-冻结,2-禁言);default:NULL;" json:"status"` // 用户的账号状态(0-封禁,1-冻结,2-禁言)
	IdAdmin       string    `gorm:"column:id_admin;type:varchar(2);comment:管理员等级（1普通,2超级）;default:NULL;" json:"id_admin"`      // 管理员等级（1普通,2超级）
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy     uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy     uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy     uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (a *AdminUser) TableName() string {
	return "admin_user"
}

type Mute struct {
	Id         uint64    `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	UserId     int64     `gorm:"column:user_id;type:varchar(50);comment:被禁言用户ID;default:NULL;" json:"user_id"`       // 被禁言用户ID
	MuteType   string    `gorm:"column:mute_type;type:varchar(20);comment:禁言范围;default:NULL;" json:"mute_type"`      // 禁言范围
	StartTime  time.Time `gorm:"column:start_time;type:datetime(3);comment:禁言开始时间;default:NULL;" json:"start_time"`  // 禁言开始时间
	EndTime    time.Time `gorm:"column:end_time;type:datetime(3);comment:禁言结束时间;default:NULL;" json:"end_time"`      // 禁言结束时间
	Reason     string    `gorm:"column:reason;type:varchar(1);comment:禁言原因;default:NULL;" json:"reason"`             // 禁言原因
	OperatorId int64     `gorm:"column:operator_id;type:bigint;comment:操作人ID;default:NULL;" json:"operator_id"`      // 操作人ID
	Status     string    `gorm:"column:status;type:varchar(10);comment:禁言状态;default:NULL;" json:"status"`            // 禁言状态
	MuteDay    string    `gorm:"column:mute_day;type:varchar(1);comment:禁言天数;default:NULL;" json:"mute_day"`         // 禁言天数
	MuteResult string    `gorm:"column:mute_result;type:varchar(50);comment:处理结果;default:NULL;" json:"mute_result"`  // 处理结果
	CreatedBy  uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy  uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy  uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (a *Mute) TableName() string {
	return "mute"
}
