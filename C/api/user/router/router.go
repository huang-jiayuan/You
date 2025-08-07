package router

import (
	"api/pkg"
	"api/user/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/sendsms", handler.Sendsms)
			user.POST("/login", handler.Login)
			user.POST("/login/password", handler.UserPassword)
			user.Use(pkg.JWTAuth("2211a"))
			user.POST("/update/password", handler.UpdatePassword)
			user.POST("/improve/message", handler.ImproveUserMessage)
			user.POST("/follow", handler.FollowUser)
			user.POST("/unfollow", handler.UnFollowUser)
		}
		// 需认证路由
		authGroup := r.Group("/auth")
		authGroup.Use(handler.AuthMiddleware())
		{

			authGroup.GET("/profile", handler.Profile)
			authGroup.POST("/logout", handler.Logout)
		}
	}

}
