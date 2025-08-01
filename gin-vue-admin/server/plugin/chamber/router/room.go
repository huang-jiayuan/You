package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Room = new(room)

type room struct {}

// Init 初始化 房间表 路由信息
func (r *room) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("room").Use(middleware.OperationRecord())
		group.POST("createRoom", apiRoom.CreateRoom)   // 新建房间表
		group.DELETE("deleteRoom", apiRoom.DeleteRoom) // 删除房间表
		group.DELETE("deleteRoomByIds", apiRoom.DeleteRoomByIds) // 批量删除房间表
		group.PUT("updateRoom", apiRoom.UpdateRoom)    // 更新房间表
	}
	{
	    group := private.Group("room")
		group.GET("findRoom", apiRoom.FindRoom)        // 根据ID获取房间表
		group.GET("getRoomList", apiRoom.GetRoomList)  // 获取房间表列表
	}
	{
	    group := public.Group("room")
	    group.GET("getRoomPublic", apiRoom.GetRoomPublic)  // 房间表开放接口
	}
}
