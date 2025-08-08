
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type RoomSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    request.PageInfo
}

// KickFromMic 踢人下麦请求结构
type KickFromMic struct {
	RoomId       int64  `json:"roomId" form:"roomId" binding:"required"`
	TargetUserId uint64 `json:"targetUserId" form:"targetUserId" binding:"required"`
	Reason       string `json:"reason" form:"reason"`
}

// MuteMicUser 禁言管理请求结构
type MuteMicUser struct {
	RoomId       int64  `json:"roomId" form:"roomId" binding:"required"`
	TargetUserId uint64 `json:"targetUserId" form:"targetUserId" binding:"required"`
	Action       int32  `json:"action" form:"action" binding:"required,oneof=1 2"` // 1-禁言, 2-解禁
	Duration     int32  `json:"duration" form:"duration"`                          // 禁言时长(分钟)
	Reason       string `json:"reason" form:"reason"`                              // 操作原因
}
