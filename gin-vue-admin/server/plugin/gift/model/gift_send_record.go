
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// GiftSendRecord giftSendRecord表 结构体
type GiftSendRecord struct {
    global.GVA_MODEL
  SendUserId  *int `json:"sendUserId" form:"sendUserId" gorm:"comment:赠送者用户ID，关联 user.id;column:send_user_id;size:20;" binding:"required"`  //赠送者用户ID，关联 user.id
  ReceiveUserId  *int `json:"receiveUserId" form:"receiveUserId" gorm:"comment:接收者用户ID，关联 user.id;column:receive_user_id;size:20;" binding:"required"`  //接收者用户ID，关联 user.id
  RoomId  *int `json:"roomId" form:"roomId" gorm:"comment:房间ID(NULL-私聊);column:room_id;size:19;" binding:"required"`  //房间ID(NULL-私聊)
  GiftId  *int `json:"giftId" form:"giftId" gorm:"comment:礼物ID;column:gift_id;size:19;" binding:"required"`  //礼物ID
  SendCount  *int `json:"sendCount" form:"sendCount" gorm:"comment:赠送数量;column:send_count;size:10;" binding:"required"`  //赠送数量
  TotalPrice  *float64 `json:"totalPrice" form:"totalPrice" gorm:"comment:总消耗(虚拟币);column:total_price;size:10;" binding:"required"`  //总消耗(虚拟币)
  TotalDiamond  *int `json:"totalDiamond" form:"totalDiamond" gorm:"comment:总消耗(钻石);column:total_diamond;size:10;" binding:"required"`  //总消耗(钻石)
  SendType  *string `json:"sendType" form:"sendType" gorm:"default:1;comment:赠送方式(1-背包,2-直接购买);column:send_type;size:1;" binding:"required"`  //赠送方式(1-背包,2-直接购买)
  Message  *string `json:"message" form:"message" gorm:"comment:赠送附言;column:message;size:255;"`  //赠送附言
  SendTime  *time.Time `json:"sendTime" form:"sendTime" gorm:"comment:赠送时间;column:send_time;"`  //赠送时间
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:状态(0-失败,1-成功,2-已撤回);column:status;size:1;"`  //状态(0-失败,1-成功,2-已撤回)
  ClientIp  *string `json:"clientIp" form:"clientIp" gorm:"comment:赠送者IP;column:client_ip;size:50;"`  //赠送者IP
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName giftSendRecord表 GiftSendRecord自定义表名 gift_send_record
func (GiftSendRecord) TableName() string {
    return "gift_send_record"
}







