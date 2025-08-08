package router

import (
	"api/room/handler"
	"github.com/gin-gonic/gin"
)
import (
	"api/pkg"
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

	api := r.Group("/api")
	{
		room := api.Group("/room")
		{
			room.Use(pkg.JWTAuth("2211a"))
			room.POST("/createroom", handler.CreateRoom)
			room.POST("/updateroom", handler.UpdateRoom)
			room.POST("/getRecommendRooms", handler.GetRecommendRooms)
			room.POST("/getRoomsByCategory", handler.GetRoomsByCategory)
			room.POST("/searchRooms", handler.SearchRooms)
			room.POST("/joinRoom", handler.JoinRoom)
			room.POST("/closeRoom", handler.CloseRoom)
			room.GET("/ws", handler.HandleWebSocket)

			// 麦位管理相关路由
			room.POST("/applyMic", handler.ApplyMic)
			room.POST("/leaveMic", handler.LeaveMic)

		}
	}
}
