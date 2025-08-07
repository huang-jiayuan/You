package hot_room

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HostRoomRouter struct {}

// InitHostRoomRouter 初始化 hostRoom表 路由信息
func (s *HostRoomRouter) InitHostRoomRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	hostRoomRouter := Router.Group("hostRoom").Use(middleware.OperationRecord())
	hostRoomRouterWithoutRecord := Router.Group("hostRoom")
	hostRoomRouterWithoutAuth := PublicRouter.Group("hostRoom")
	{
		hostRoomRouter.POST("createHostRoom", hostRoomApi.CreateHostRoom)   // 新建hostRoom表
		hostRoomRouter.DELETE("deleteHostRoom", hostRoomApi.DeleteHostRoom) // 删除hostRoom表
		hostRoomRouter.DELETE("deleteHostRoomByIds", hostRoomApi.DeleteHostRoomByIds) // 批量删除hostRoom表
		hostRoomRouter.PUT("updateHostRoom", hostRoomApi.UpdateHostRoom)    // 更新hostRoom表
	}
	{
		hostRoomRouterWithoutRecord.GET("findHostRoom", hostRoomApi.FindHostRoom)        // 根据ID获取hostRoom表
		hostRoomRouterWithoutRecord.GET("getHostRoomList", hostRoomApi.GetHostRoomList)  // 获取hostRoom表列表
	}
	{
	    hostRoomRouterWithoutAuth.GET("getHostRoomPublic", hostRoomApi.GetHostRoomPublic)  // hostRoom表开放接口
	}
}
