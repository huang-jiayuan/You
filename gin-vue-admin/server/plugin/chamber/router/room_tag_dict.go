package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var RoomTagDict = new(roomTagDict)

type roomTagDict struct {}

// Init 初始化 房间话题表 路由信息
func (r *roomTagDict) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("roomTagDict").Use(middleware.OperationRecord())
		group.POST("createRoomTagDict", apiRoomTagDict.CreateRoomTagDict)   // 新建房间话题表
		group.DELETE("deleteRoomTagDict", apiRoomTagDict.DeleteRoomTagDict) // 删除房间话题表
		group.DELETE("deleteRoomTagDictByIds", apiRoomTagDict.DeleteRoomTagDictByIds) // 批量删除房间话题表
		group.PUT("updateRoomTagDict", apiRoomTagDict.UpdateRoomTagDict)    // 更新房间话题表
	}
	{
	    group := private.Group("roomTagDict")
		group.GET("findRoomTagDict", apiRoomTagDict.FindRoomTagDict)        // 根据ID获取房间话题表
		group.GET("getRoomTagDictList", apiRoomTagDict.GetRoomTagDictList)  // 获取房间话题表列表
	}
	{
	    group := public.Group("roomTagDict")
	    group.GET("getRoomTagDictPublic", apiRoomTagDict.GetRoomTagDictPublic)  // 房间话题表开放接口
	}
}
