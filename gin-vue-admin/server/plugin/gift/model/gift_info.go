
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// GiftInfo giftInfo表 结构体
type GiftInfo struct {
    global.GVA_MODEL
  GiftId  *int `json:"giftId" form:"giftId" gorm:"primarykey;comment:礼物唯一ID;column:gift_id;size:19;"`  //礼物唯一ID
  GiftName  *string `json:"giftName" form:"giftName" gorm:"comment:礼物名称;column:gift_name;size:100;" binding:"required"`  //礼物名称
  GiftType  *string `json:"giftType" form:"giftType" gorm:"default:0;comment:礼物类型(1-普通,2-特效,3-稀有,4-自定义);column:gift_type;size:1;" binding:"required"`  //礼物类型(1-普通,2-特效,3-稀有,4-自定义)
  Price  *float64 `json:"price" form:"price" gorm:"comment:虚拟币价格;column:price;"`  //虚拟币价格
  DiamondPrice  *int `json:"diamondPrice" form:"diamondPrice" gorm:"comment:钻石价格(可为空);column:diamond_price;"`  //钻石价格(可为空)
  IconUrl  *string `json:"iconUrl" form:"iconUrl" gorm:"comment:礼物图标URL;column:icon_url;size:255;"`  //礼物图标URL
  AnimationUrl  *string `json:"animationUrl" form:"animationUrl" gorm:"comment:动画URL(特效礼物);column:animation_url;size:255;"`  //动画URL(特效礼物)
  Weight  *int `json:"weight" form:"weight" gorm:"comment:排序权重;column:weight;"`  //排序权重
  IsHot  *string `json:"isHot" form:"isHot" gorm:"default:0;comment:是否热门(1-是,0-否);column:is_hot;size:1;"`  //是否热门(1-是,0-否)
  IsLimit  *string `json:"isLimit" form:"isLimit" gorm:"default:0;comment:是否限时(1-是,0-否);column:is_limit;size:1;"`  //是否限时(1-是,0-否)
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:状态(0-下架,1-上架,2-待审核);column:status;size:1;" binding:"required"`  //状态(0-下架,1-上架,2-待审核)
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName giftInfo表 GiftInfo自定义表名 gift_info
func (GiftInfo) TableName() string {
    return "gift_info"
}







