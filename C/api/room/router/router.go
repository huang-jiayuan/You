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
			// 所有接口都不需要JWT验证，方便开发调试
			room.POST("/getRecommendRooms", handler.GetRecommendRooms)
			room.POST("/getRoomsByCategory", handler.GetRoomsByCategory)
			room.POST("/searchRooms", handler.SearchRooms)
			room.POST("/joinRoom", handler.JoinRoom)
			room.POST("/createroom", handler.CreateRoom)
			room.POST("/sendgift", handler.SendGifts)
			room.POST("/kick", handler.KickUser)
			room.POST("/mute", handler.MuteUser)
			room.POST("/unmute", handler.UnmuteUser)
			room.POST("/updateroom", handler.UpdateRoom)
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
