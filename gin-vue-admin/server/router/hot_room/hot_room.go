package hot_room

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HotRoomRouter struct {}

// InitHotRoomRouter 初始化 hotRoom表 路由信息
func (s *HotRoomRouter) InitHotRoomRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	hotRoomRouter := Router.Group("hotRoom").Use(middleware.OperationRecord())
	hotRoomRouterWithoutRecord := Router.Group("hotRoom")
	hotRoomRouterWithoutAuth := PublicRouter.Group("hotRoom")
	{
		hotRoomRouter.POST("createHotRoom", hotRoomApi.CreateHotRoom)   // 新建hotRoom表
		hotRoomRouter.DELETE("deleteHotRoom", hotRoomApi.DeleteHotRoom) // 删除hotRoom表
		hotRoomRouter.DELETE("deleteHotRoomByIds", hotRoomApi.DeleteHotRoomByIds) // 批量删除hotRoom表
		hotRoomRouter.PUT("updateHotRoom", hotRoomApi.UpdateHotRoom)    // 更新hotRoom表
	}
	{
		hotRoomRouterWithoutRecord.GET("findHotRoom", hotRoomApi.FindHotRoom)        // 根据ID获取hotRoom表
		hotRoomRouterWithoutRecord.GET("getHotRoomList", hotRoomApi.GetHotRoomList)  // 获取hotRoom表列表
	}
	{
	    hotRoomRouterWithoutAuth.GET("getHotRoomPublic", hotRoomApi.GetHotRoomPublic)  // hotRoom表开放接口
	}
}
