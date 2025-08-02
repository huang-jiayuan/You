package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Mute = new(mute)

type mute struct {}

// Init 初始化 mute表 路由信息
func (r *mute) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("mute").Use(middleware.OperationRecord())
		group.POST("createMute", apiMute.CreateMute)   // 新建mute表
		group.DELETE("deleteMute", apiMute.DeleteMute) // 删除mute表
		group.DELETE("deleteMuteByIds", apiMute.DeleteMuteByIds) // 批量删除mute表
		group.PUT("updateMute", apiMute.UpdateMute)    // 更新mute表
	}
	{
	    group := private.Group("mute")
		group.GET("findMute", apiMute.FindMute)        // 根据ID获取mute表
		group.GET("getMuteList", apiMute.GetMuteList)  // 获取mute表列表
	}
	{
	    group := public.Group("mute")
	    group.GET("getMutePublic", apiMute.GetMutePublic)  // mute表开放接口
	}
}
