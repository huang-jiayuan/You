
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HostRoomSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      RoomId  *int `json:"roomId" form:"roomId"` 
      RoomType  *string `json:"roomType" form:"roomType"` 
      RoomTags  *string `json:"roomTags" form:"roomTags"` 
      RoomStatus  *string `json:"roomStatus" form:"roomStatus"` 
      RoomHost  *string `json:"roomHost" form:"roomHost"` 
      RoomNum  *int `json:"roomNum" form:"roomNum"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
