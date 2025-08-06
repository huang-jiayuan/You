
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Mute mute表 结构体
type Mute struct {
    global.GVA_MODEL
  UserId  *string `json:"userId" form:"userId" gorm:"comment:被禁言用户ID;column:user_id;size:50;" binding:"required"`  //被禁言用户ID
  MuteType  *string `json:"muteType" form:"muteType" gorm:"comment:禁言范围;column:mute_type;size:20;" binding:"required"`  //禁言范围
  StartTime  *time.Time `json:"startTime" form:"startTime" gorm:"comment:禁言开始时间;column:start_time;" binding:"required"`  //禁言开始时间
  EndTime  *time.Time `json:"endTime" form:"endTime" gorm:"comment:禁言结束时间;column:end_time;" binding:"required"`  //禁言结束时间
  Reason  *string `json:"reason" form:"reason" gorm:"comment:禁言原因;column:reason;size:1;" binding:"required"`  //禁言原因
  OperatorId  *int `json:"operatorId" form:"operatorId" gorm:"comment:操作人ID;column:operator_id;size:10;" binding:"required"`  //操作人ID
  Status  *string `json:"status" form:"status" gorm:"comment:禁言状态;column:status;size:10;" binding:"required"`  //禁言状态
  MuteDay  *string `json:"muteDay" form:"muteDay" gorm:"comment:禁言天数;column:mute_day;size:1;" binding:"required"`  //禁言天数
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName mute表 Mute自定义表名 mute
func (Mute) TableName() string {
    return "mute"
}







