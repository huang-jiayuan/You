package router

import (
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
		}
	}

}
