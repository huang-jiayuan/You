package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var UserGiftBackpack = new(userGiftBackpack)

type userGiftBackpack struct {}

// Init 初始化 用户背包礼物表 路由信息
func (r *userGiftBackpack) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("userGiftBackpack").Use(middleware.OperationRecord())
		group.POST("createUserGiftBackpack", apiUserGiftBackpack.CreateUserGiftBackpack)   // 新建用户背包礼物表
		group.DELETE("deleteUserGiftBackpack", apiUserGiftBackpack.DeleteUserGiftBackpack) // 删除用户背包礼物表
		group.DELETE("deleteUserGiftBackpackByIds", apiUserGiftBackpack.DeleteUserGiftBackpackByIds) // 批量删除用户背包礼物表
		group.PUT("updateUserGiftBackpack", apiUserGiftBackpack.UpdateUserGiftBackpack)    // 更新用户背包礼物表
	}
	{
	    group := private.Group("userGiftBackpack")
		group.GET("findUserGiftBackpack", apiUserGiftBackpack.FindUserGiftBackpack)        // 根据ID获取用户背包礼物表
		group.GET("getUserGiftBackpackList", apiUserGiftBackpack.GetUserGiftBackpackList)  // 获取用户背包礼物表列表
	}
	{
	    group := public.Group("userGiftBackpack")
	    group.GET("getUserGiftBackpackPublic", apiUserGiftBackpack.GetUserGiftBackpackPublic)  // 用户背包礼物表开放接口
	}
}
