
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Room 房间表 结构体
type Room struct {
    global.GVA_MODEL
  RoomName  *string `json:"roomName" form:"roomName" gorm:"comment:房间名称;column:room_name;size:50;" binding:"required"`  //房间名称
  UserId  *int `json:"userId" form:"userId" gorm:"comment:房主用户ID，关联 user.id;column:user_id;size:20;" binding:"required"`  //房主用户ID，关联 user.id
  RoomType  *string `json:"roomType" form:"roomType" gorm:"default:0;comment:房间类型(0-公开房,1-私密房);column:room_type;size:1;" binding:"required"`  //房间类型(0-公开房,1-私密房)
  Status  *string `json:"status" form:"status" gorm:"default:0;comment:房间状态(0-正常,1-已解散,2-全局禁言);column:status;size:1;" binding:"required"`  //房间状态(0-正常,1-已解散,2-全局禁言)
  Announcement  *string `json:"announcement" form:"announcement" gorm:"comment:房间公告;column:announcement;size:200;"`  //房间公告
  ClosedAt  *time.Time `json:"closedAt" form:"closedAt" gorm:"comment:解散时间;column:closed_at;"`  //解散时间
  LastActiveTime  *time.Time `json:"lastActiveTime" form:"lastActiveTime" gorm:"comment:最后活跃时间;column:last_active_time;"`  //最后活跃时间
  FkMemberRoom  *int `json:"fkMemberRoom" form:"fkMemberRoom" gorm:"comment:房间人数;column:fk_member_room;size:10;"`  //房间人数
  Image  *string `json:"image" form:"image" gorm:"comment:房间封面图;column:image;size:255;"`  //房间封面图
  TagId  *int `json:"tagId" form:"tagId" gorm:"comment:关联的标签ID（一个房间只能选一个）;column:tag_id;size:20;"`  //关联的标签ID（一个房间只能选一个）
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 房间表 Room自定义表名 room
func (Room) TableName() string {
    return "room"
}







