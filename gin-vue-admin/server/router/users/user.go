package users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {}

// InitUserRouter 初始化 user表 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	user1Router := Router.Group("user1").Use(middleware.OperationRecord())
	user1RouterWithoutRecord := Router.Group("user1")
	user1RouterWithoutAuth := PublicRouter.Group("user1")
	{
		user1Router.POST("createUser", user1Api.CreateUser)   // 新建user表
		user1Router.DELETE("deleteUser", user1Api.DeleteUser) // 删除user表
		user1Router.DELETE("deleteUserByIds", user1Api.DeleteUserByIds) // 批量删除user表
		user1Router.PUT("updateUser", user1Api.UpdateUser)    // 更新user表
	}
	{
		user1RouterWithoutRecord.GET("findUser", user1Api.FindUser)        // 根据ID获取user表
		user1RouterWithoutRecord.GET("getUserList", user1Api.GetUserList)  // 获取user表列表
	}
	{
	    user1RouterWithoutAuth.GET("getUserPublic", user1Api.GetUserPublic)  // user表开放接口
	}
}
