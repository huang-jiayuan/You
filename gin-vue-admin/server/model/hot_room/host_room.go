
// 自动生成模板HostRoom
package hot_room
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// hostRoom表 结构体  HostRoom
type HostRoom struct {
    global.GVA_MODEL
  RoomId  *int `json:"roomId" form:"roomId" gorm:"comment:房间id;column:room_id;size:10;" binding:"required"`  //房间id
  RoomType  *string `json:"roomType" form:"roomType" gorm:"comment:房间类型;column:room_type;size:50;" binding:"required"`  //房间类型
  RoomTags  *string `json:"roomTags" form:"roomTags" gorm:"comment:房间标签;column:room_tags;size:50;" binding:"required"`  //房间标签
  RoomStatus  *string `json:"roomStatus" form:"roomStatus" gorm:"comment:房间状态;column:room_status;size:50;" binding:"required"`  //房间状态
  RoomHost  *string `json:"roomHost" form:"roomHost" gorm:"comment:房主;column:room_host;size:50;" binding:"required"`  //房主
  RoomNum  *int `json:"roomNum" form:"roomNum" gorm:"comment:房间人数;column:room_num;size:10;" binding:"required"`  //房间人数
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName hostRoom表 HostRoom自定义表名 host_room
func (HostRoom) TableName() string {
    return "host_room"
}





