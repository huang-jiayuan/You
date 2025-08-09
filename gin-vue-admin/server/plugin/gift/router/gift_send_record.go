package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var GiftSendRecord = new(giftSendRecord)

type giftSendRecord struct {}

// Init 初始化 giftSendRecord表 路由信息
func (r *giftSendRecord) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("giftSendRecord").Use(middleware.OperationRecord())
		group.POST("createGiftSendRecord", apiGiftSendRecord.CreateGiftSendRecord)   // 新建giftSendRecord表
		group.DELETE("deleteGiftSendRecord", apiGiftSendRecord.DeleteGiftSendRecord) // 删除giftSendRecord表
		group.DELETE("deleteGiftSendRecordByIds", apiGiftSendRecord.DeleteGiftSendRecordByIds) // 批量删除giftSendRecord表
		group.PUT("updateGiftSendRecord", apiGiftSendRecord.UpdateGiftSendRecord)    // 更新giftSendRecord表
	}
	{
	    group := private.Group("giftSendRecord")
		group.GET("findGiftSendRecord", apiGiftSendRecord.FindGiftSendRecord)        // 根据ID获取giftSendRecord表
		group.GET("getGiftSendRecordList", apiGiftSendRecord.GetGiftSendRecordList)  // 获取giftSendRecord表列表
	}
	{
	    group := public.Group("giftSendRecord")
	    group.GET("getGiftSendRecordPublic", apiGiftSendRecord.GetGiftSendRecordPublic)  // giftSendRecord表开放接口
	}
}
