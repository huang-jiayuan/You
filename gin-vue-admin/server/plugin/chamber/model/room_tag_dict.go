
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RoomTagDict 房间话题表 结构体
type RoomTagDict struct {
    global.GVA_MODEL
  TagName  *string `json:"tagName" form:"tagName" gorm:"comment:标签名称;column:tag_name;size:30;" binding:"required"`  //标签名称
  TagDesc  *string `json:"tagDesc" form:"tagDesc" gorm:"comment:标签描述;column:tag_desc;size:100;" binding:"required"`  //标签描述
  Status  *string `json:"status" form:"status" gorm:"comment:状态(0-启用,1-停用);column:status;size:1;" binding:"required"`  //状态(0-启用,1-停用)
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 房间话题表 RoomTagDict自定义表名 room_tag_dict
func (RoomTagDict) TableName() string {
    return "room_tag_dict"
}







