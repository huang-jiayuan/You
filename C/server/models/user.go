package models

import (
	"server/pkg"
	"server/user/basic/global"
	__ "server/user/proto"
	"time"
)

type User struct {
	Id            uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:用户唯一ID;primaryKey;" json:"id"`                            // 用户唯一ID
	Username      string    `gorm:"column:username;type:varchar(32);comment:登录账号(手机号/第三方账号);not null;" json:"username"`             // 登录账号(手机号/第三方账号)
	Password      string    `gorm:"column:password;type:varchar(64);comment:登录密码(SHA-256加密);" json:"password"`                      // 登录密码(SHA-256加密)
	Nickname      string    `gorm:"column:nickname;type:varchar(32);comment:用户昵称;not null;" json:"nickname"`                        // 用户昵称
	Avatar        string    `gorm:"column:avatar;type:varchar(255);comment:头像URL;" json:"avatar"`                                   // 头像URL
	Gender        string    `gorm:"column:gender;type:varchar(50);comment:性别(0-保密,1-男,2-女);" json:"gender"`                         // 性别(0-保密,1-男,2-女)
	VoiceTag      string    `gorm:"column:voice_tag;type:varchar(100);comment:声音标签(逗号分隔);" json:"voice_tag"`                        // 声音标签(逗号分隔)
	InterestTags  string    `gorm:"column:interest_tags;type:varchar(200);comment:兴趣标签(逗号分隔);" json:"interest_tags"`                // 兴趣标签(逗号分隔)
	Level         int16     `gorm:"column:level;type:smallint;comment:用户等级;" json:"level"`                                          // 用户等级
	VipStatus     string    `gorm:"column:vip_status;type:varchar(30);comment:会员状态(0-非会员,1-月会员,2-年会员);" json:"vip_status"`          // 会员状态(0-非会员,1-月会员,2-年会员)
	VipExpireTime time.Time `gorm:"column:vip_expire_time;type:datetime(3);default:null;comment:会员过期时间;" json:"vip_expire_time"`    // 会员过期时间
	Balance       float64   `gorm:"column:balance;type:float;comment:虚拟币余额;" json:"balance"`                                        // 虚拟币余额
	Diamond       int16     `gorm:"column:diamond;type:smallint;comment:钻石数量;" json:"diamond"`                                      // 钻石数量
	RegisterTime  time.Time `gorm:"column:register_time;type:datetime(3);comment:注册时间;not null;" json:"register_time"`              // 注册时间
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime(3);comment:最后登录时间;" json:"last_login_time"`                 // 最后登录时间
	LastLoginIp   string    `gorm:"column:last_login_ip;type:varchar(50);comment:最后登录IP;not null;" json:"last_login_ip"`            // 最后登录IP
	Status        string    `gorm:"column:status;type:varchar(30);comment:账号状态(0-正常,1-封禁,2-冻结,3-禁言);" json:"status"`                // 账号状态(0-正常,1-封禁,2-冻结,3-禁言)
	FreezeReason  string    `gorm:"column:freeze_reason;type:varchar(200);comment:冻结/封禁原因;" json:"freeze_reason"`                   // 冻结/封禁原因
	FreezeEndTime time.Time `gorm:"column:freeze_end_time;type:datetime(3);default:null;comment:冻结/封禁结束时间;" json:"freeze_end_time"` // 冻结/封禁结束时间
	IsAuth        string    `gorm:"column:is_auth;type:varchar(30);comment:实名认证状态（0：未认证1：已认证2）;" json:"is_auth"`                    // 实名认证状态（0：未认证1：已认证2）
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime(3);default:null;" json:"deleted_at"`
	CreatedBy     uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;" json:"created_by"` // 创建者
	UpdatedBy     uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;" json:"updated_by"` // 更新者
	DeletedBy     uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;" json:"deleted_by"` // 删除者
}

func (u *User) TableName() string {
	return "user"
}

// todo:根据手机号查询用户信息
func (u *User) FindUserByMobile(phone string) (*User, error) {
	err := global.DB.Where("username=?", phone).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// todo:用户注册
func (u *User) CreateUser(phone, lastLoginIp string) (*User, error) {
	user := &User{
		Username:      phone,
		Password:      pkg.PwdMd5("123456"), //默认密码
		RegisterTime:  time.Now(),
		LastLoginIp:   lastLoginIp,
		LastLoginTime: time.Now(),
	}
	err := global.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) UpdatePassword(id int64, password string) error {
	return global.DB.Model(&u).Where("id=?", id).Update("password", pkg.PwdMd5(password)).Error
}

// 根据用户id查询用户信息
func (u *User) FindUserById(id int64) (*User, error) {
	err := global.DB.Where("id=?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) ImproveUserMessage(in *__.ImproveUserMessageRequest) error {
	return global.DB.Model(&u).Where("id=?", in.UserId).Updates(&User{
		Nickname:     in.Nickname,
		Avatar:       in.Avatar,
		Gender:       in.Gender,
		VoiceTag:     in.VoiceTag,
		InterestTags: in.InterestTags,
	}).Error
}
