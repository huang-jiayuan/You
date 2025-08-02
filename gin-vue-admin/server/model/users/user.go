
// 自动生成模板User
package users
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// user表 结构体  User
type User struct {
    global.GVA_MODEL
  Username  *string `json:"username" form:"username" gorm:"comment:登录账号(手机号/第三方账号);column:username;size:32;" binding:"required"`  //登录账号(手机号/第三方账号)
  Password  *string `json:"password" form:"password" gorm:"comment:登录密码(SHA-256加密);column:password;size:64;" binding:"required"`  //登录密码(SHA-256加密)
  Nickname  *string `json:"nickname" form:"nickname" gorm:"comment:用户昵称;column:nickname;size:32;" binding:"required"`  //用户昵称
  Avatar  *string `json:"avatar" form:"avatar" gorm:"comment:头像URL;column:avatar;size:255;" binding:"required"`  //头像URL
  Gender  *string `json:"gender" form:"gender" gorm:"comment:性别(0-保密,1-男,2-女);column:gender;size:50;" binding:"required"`  //性别(0-保密,1-男,2-女)
  VoiceTag  *string `json:"voiceTag" form:"voiceTag" gorm:"comment:声音标签(逗号分隔);column:voice_tag;size:100;"`  //声音标签(逗号分隔)
  InterestTags  *string `json:"interestTags" form:"interestTags" gorm:"comment:兴趣标签(逗号分隔);column:interest_tags;size:200;"`  //兴趣标签(逗号分隔)
  Level  *int `json:"level" form:"level" gorm:"comment:用户等级;column:level;size:10;" binding:"required"`  //用户等级
  VipStatus  *string `json:"vipStatus" form:"vipStatus" gorm:"comment:会员状态(0-非会员,1-月会员,2-年会员);column:vip_status;size:30;" binding:"required"`  //会员状态(0-非会员,1-月会员,2-年会员)
  VipExpireTime  *time.Time `json:"vipExpireTime" form:"vipExpireTime" gorm:"comment:会员过期时间;column:vip_expire_time;"`  //会员过期时间
  Balance  *float64 `json:"balance" form:"balance" gorm:"comment:虚拟币余额;column:balance;size:10;" binding:"required"`  //虚拟币余额
  Diamond  *int `json:"diamond" form:"diamond" gorm:"comment:钻石数量;column:diamond;size:10;" binding:"required"`  //钻石数量
  RegisterTime  *time.Time `json:"registerTime" form:"registerTime" gorm:"comment:注册时间;column:register_time;" binding:"required"`  //注册时间
  LastLoginTime  *time.Time `json:"lastLoginTime" form:"lastLoginTime" gorm:"comment:最后登录时间;column:last_login_time;" binding:"required"`  //最后登录时间
  LastLoginIp  *string `json:"lastLoginIp" form:"lastLoginIp" gorm:"comment:最后登录IP;column:last_login_ip;size:50;" binding:"required"`  //最后登录IP
  Status  *string `json:"status" form:"status" gorm:"comment:账号状态(0-正常,1-封禁,2-冻结,3-禁言);column:status;size:30;" binding:"required"`  //账号状态(0-正常,1-封禁,2-冻结,3-禁言)
  FreezeReason  *string `json:"freezeReason" form:"freezeReason" gorm:"comment:冻结/封禁原因;column:freeze_reason;size:200;"`  //冻结/封禁原因
  FreezeEndTime  *time.Time `json:"freezeEndTime" form:"freezeEndTime" gorm:"comment:冻结/封禁结束时间;column:freeze_end_time;"`  //冻结/封禁结束时间
  IsAuth  *string `json:"isAuth" form:"isAuth" gorm:"comment:实名认证状态（0：未认证1：已认证2）;column:is_auth;size:30;" binding:"required"`  //实名认证状态（0：未认证1：已认证2）
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName user表 User自定义表名 user
func (User) TableName() string {
    return "user"
}





