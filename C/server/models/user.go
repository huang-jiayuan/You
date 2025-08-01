package models

import (
	"server/user/basic/global"
	"time"
)

type User struct {
	Id            uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:用户唯一ID;primaryKey;" json:"id"`                                              // 用户唯一ID
	Username      string    `gorm:"column:username;type:varchar(32);comment:登录账号(手机号/第三方账号);not null;" json:"username"`                       // 登录账号(手机号/第三方账号)
	Password      string    `gorm:"column:password;type:varchar(64);comment:登录密码(SHA-256加密);" json:"password"`                                      // 登录密码(SHA-256加密)
	Nickname      string    `gorm:"column:nickname;type:varchar(32);comment:用户昵称;not null;" json:"nickname"`                                          // 用户昵称
	Avatar        string    `gorm:"column:avatar;type:varchar(255);comment:头像URL;not null;default:/static/avatar/default.png;" json:"avatar"`           // 头像URL
	Gender        int8      `gorm:"column:gender;type:tinyint;comment:性别(0-保密,1-男,2-女);not null;default:0;" json:"gender"`                          // 性别(0-保密,1-男,2-女)
	VoiceTag      string    `gorm:"column:voice_tag;type:varchar(100);comment:声音标签(逗号分隔);" json:"voice_tag"`                                      // 声音标签(逗号分隔)
	InterestTags  string    `gorm:"column:interest_tags;type:varchar(200);comment:兴趣标签(逗号分隔);" json:"interest_tags"`                              // 兴趣标签(逗号分隔)
	Level         int32     `gorm:"column:level;type:int;comment:用户等级;not null;default:1;" json:"level"`                                              // 用户等级
	VipStatus     int8      `gorm:"column:vip_status;type:tinyint;comment:会员状态(0-非会员,1-月会员,2-年会员);not null;default:0;" json:"vip_status"`    // 会员状态(0-非会员,1-月会员,2-年会员)
	VipExpireTime time.Time `gorm:"column:vip_expire_time;type:datetime;comment:会员过期时间;" json:"vip_expire_time"`                                    // 会员过期时间
	Balance       float64   `gorm:"column:balance;type:decimal(10, 2);comment:虚拟币余额;not null;default:0.00;" json:"balance"`                          // 虚拟币余额
	Diamond       int32     `gorm:"column:diamond;type:int;comment:钻石数量;not null;default:0;" json:"diamond"`                                          // 钻石数量
	RegisterTime  time.Time `gorm:"column:register_time;type:datetime(3);comment:注册时间;not null;" json:"register_time"`                                // 注册时间
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime;comment:最后登录时间;not null;default:CURRENT_TIMESTAMP;" json:"last_login_time"` // 最后登录时间
	LastLoginIp   string    `gorm:"column:last_login_ip;type:varchar(50);comment:最后登录IP;not null;" json:"last_login_ip"`                              // 最后登录IP
	Status        int8      `gorm:"column:status;type:tinyint;comment:账号状态(0-正常,1-封禁,2-冻结);not null;default:0;" json:"status"`                  // 账号状态(0-正常,1-封禁,2-冻结)
	FreezeReason  string    `gorm:"column:freeze_reason;type:varchar(200);comment:冻结/封禁原因;" json:"freeze_reason"`                                   // 冻结/封禁原因
	FreezeEndTime time.Time `gorm:"column:freeze_end_time;type:datetime;comment:冻结/封禁结束时间;" json:"freeze_end_time"`                               // 冻结/封禁结束时间
	IsAuth        uint32    `gorm:"column:is_auth;type:int UNSIGNED ZEROFILL;comment:实名认证状态（0：未认证1：已认证2）;not null;default:0;" json:"is_auth"` // 实名认证状态（0：未认证1：已认证2）
}

func (u *User) TableName() string {
	return "user"
}

// todo:根据手机号查询用户信息
func (u *User) FindUserByMobile(phone string) (*User, error) {
	err := global.DB.Where("phone=?", phone).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// todo:用户注册
func (u *User) CreateUser(phone string) (*User, error) {
	user := &User{
		Username: phone,
	}
	err := global.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
