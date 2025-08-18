package router

import (
	"api/pkg"
	"api/user/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"message":   "Backend service is running",
			"port":      "8081",
			"timestamp": c.GetHeader("Date"),
		})
	})

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
			user.POST("/follow/list", handler.UserFollowList)
			user.POST("/center", handler.UserCenterList)
			// 在Gin路由中添加WebSocket端点
			user.GET("/ws", handler.HandleWebSocket)
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
