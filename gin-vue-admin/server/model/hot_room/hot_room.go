
// 自动生成模板HotRoom
package hot_room
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// hotRoom表 结构体  HotRoom
type HotRoom struct {
    global.GVA_MODEL
  RoomId  *int `json:"roomId" form:"roomId" gorm:"comment:房间id;column:room_id;size:19;" binding:"required"`  //房间id
  RoomType  *string `json:"roomType" form:"roomType" gorm:"comment:房间类型;column:room_type;size:30;" binding:"required"`  //房间类型
  RoomTags  *string `json:"roomTags" form:"roomTags" gorm:"comment:房间标签;column:room_tags;size:60;" binding:"required"`  //房间标签
  RoomStatus  *string `json:"roomStatus" form:"roomStatus" gorm:"comment:房间状态;column:room_status;size:30;" binding:"required"`  //房间状态
  RoomHost  *string `json:"roomHost" form:"roomHost" gorm:"comment:房主;column:room_host;size:20;" binding:"required"`  //房主
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName hotRoom表 HotRoom自定义表名 hot_room
func (HotRoom) TableName() string {
    return "hot_room"
}





