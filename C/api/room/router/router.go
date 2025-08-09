package router

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
<<<<<<< HEAD
			room.POST("/sendgift", handler.SendGifts)
			room.Use(pkg.JWTAuth("2211a"))
			room.POST("/setadmin", handler.SetAdmin)
=======

			room.Use(pkg.JWTAuth("2211a"))
			room.POST("/sendgift", handler.SendGifts)
			room.POST("/kick", handler.KickUser)
			room.POST("/mute", handler.MuteUser)
			room.POST("/unmute", handler.UnmuteUser)
>>>>>>> c43f73117726d705c0d478e5c7124ee2a97e67b5
			room.POST("/createroom", handler.CreateRoom)
			room.POST("/updateroom", handler.UpdateRoom)
			room.POST("/getRecommendRooms", handler.GetRecommendRooms)
			room.POST("/getRoomsByCategory", handler.GetRoomsByCategory)
			room.POST("/searchRooms", handler.SearchRooms)
			room.POST("/joinRoom", handler.JoinRoom)
			room.POST("/closeRoom", handler.CloseRoom)
			room.GET("/ws", handler.HandleWebSocket)
			room.POST("/setadmin", handler.SetAdmin)
			// 麦位管理相关路由
			room.POST("/applyMic", handler.ApplyMic)
			room.POST("/leaveMic", handler.LeaveMic)

			room.POST("/handleMicApplication", handler.HandleMicApplication)
			room.POST("/kickFromMic", handler.KickFromMic)
			room.POST("/muteMicUser", handler.MuteMicUser)
		}
	}
}
