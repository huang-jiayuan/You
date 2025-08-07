package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var AdminUser = new(adminUser)

type adminUser struct {}

// Init 初始化 管理员表 路由信息
func (r *adminUser) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("adminUser").Use(middleware.OperationRecord())
		group.POST("createAdminUser", apiAdminUser.CreateAdminUser)   // 新建管理员表
		group.DELETE("deleteAdminUser", apiAdminUser.DeleteAdminUser) // 删除管理员表
		group.DELETE("deleteAdminUserByIds", apiAdminUser.DeleteAdminUserByIds) // 批量删除管理员表
		group.PUT("updateAdminUser", apiAdminUser.UpdateAdminUser)    // 更新管理员表
	}
	{
	    group := private.Group("adminUser")
		group.GET("findAdminUser", apiAdminUser.FindAdminUser)        // 根据ID获取管理员表
		group.GET("getAdminUserList", apiAdminUser.GetAdminUserList)  // 获取管理员表列表
	}
	{
	    group := public.Group("adminUser")
	    group.GET("getAdminUserPublic", apiAdminUser.GetAdminUserPublic)  // 管理员表开放接口
	}
}
