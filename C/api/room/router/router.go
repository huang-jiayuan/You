package router

import (
	"api/room/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	api := r.Group("/api")
	{
		room := api.Group("/room")
		{
			room.POST("/sendgift", handler.SendGifts)
		}
	}

}
