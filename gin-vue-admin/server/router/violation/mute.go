package violation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MuteRouter struct {}

// InitMuteRouter 初始化 mute表 路由信息
func (s *MuteRouter) InitMuteRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	muteRouter := Router.Group("mute").Use(middleware.OperationRecord())
	muteRouterWithoutRecord := Router.Group("mute")
	muteRouterWithoutAuth := PublicRouter.Group("mute")
	{
		muteRouter.POST("createMute", muteApi.CreateMute)   // 新建mute表
		muteRouter.DELETE("deleteMute", muteApi.DeleteMute) // 删除mute表
		muteRouter.DELETE("deleteMuteByIds", muteApi.DeleteMuteByIds) // 批量删除mute表
		muteRouter.PUT("updateMute", muteApi.UpdateMute)    // 更新mute表
	}
	{
		muteRouterWithoutRecord.GET("findMute", muteApi.FindMute)        // 根据ID获取mute表
		muteRouterWithoutRecord.GET("getMuteList", muteApi.GetMuteList)  // 获取mute表列表
	}
	{
	    muteRouterWithoutAuth.GET("getMutePublic", muteApi.GetMutePublic)  // mute表开放接口
	}
}
