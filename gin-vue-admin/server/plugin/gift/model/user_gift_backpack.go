
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// UserGiftBackpack 用户背包礼物表 结构体
type UserGiftBackpack struct {
    global.GVA_MODEL
  UserId  *int `json:"userId" form:"userId" gorm:"comment:所属用户ID，关联 user.id;column:user_id;size:20;" binding:"required"`  //所属用户ID，关联 user.id
  GiftId  *int `json:"giftId" form:"giftId" gorm:"comment:礼物ID;column:gift_id;size:19;" binding:"required"`  //礼物ID
  Count  *int `json:"count" form:"count" gorm:"comment:持有数量;column:count;size:10;"`  //持有数量
  ObtainSource  *string `json:"obtainSource" form:"obtainSource" gorm:"default:0;comment:获得来源;column:obtain_source;size:50;"`  //获得来源
  ExpireTime  *time.Time `json:"expireTime" form:"expireTime" gorm:"comment:过期时间(NULL-永久);column:expire_time;"`  //过期时间(NULL-永久)
  IsValid  *string `json:"isValid" form:"isValid" gorm:"default:0;comment:是否有效(1-是,0-否);column:is_valid;size:1;"`  //是否有效(1-是,0-否)
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 用户背包礼物表 UserGiftBackpack自定义表名 user_gift_backpack
func (UserGiftBackpack) TableName() string {
    return "user_gift_backpack"
}







