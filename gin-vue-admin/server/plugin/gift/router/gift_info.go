package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var GiftInfo = new(giftInfo)

type giftInfo struct {}

// Init 初始化 giftInfo表 路由信息
func (r *giftInfo) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("giftInfo").Use(middleware.OperationRecord())
		group.POST("createGiftInfo", apiGiftInfo.CreateGiftInfo)   // 新建giftInfo表
		group.DELETE("deleteGiftInfo", apiGiftInfo.DeleteGiftInfo) // 删除giftInfo表
		group.DELETE("deleteGiftInfoByIds", apiGiftInfo.DeleteGiftInfoByIds) // 批量删除giftInfo表
		group.PUT("updateGiftInfo", apiGiftInfo.UpdateGiftInfo)    // 更新giftInfo表
	}
	{
	    group := private.Group("giftInfo")
		group.GET("findGiftInfo", apiGiftInfo.FindGiftInfo)        // 根据ID获取giftInfo表
		group.GET("getGiftInfoList", apiGiftInfo.GetGiftInfoList)  // 获取giftInfo表列表
	}
	{
	    group := public.Group("giftInfo")
	    group.GET("getGiftInfoPublic", apiGiftInfo.GetGiftInfoPublic)  // giftInfo表开放接口
	}
}
