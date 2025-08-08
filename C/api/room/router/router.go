package router

import (
<<<<<<< HEAD
=======
	"api/pkg"
>>>>>>> c0418eef7b50d6eb94c6431ae71816618e38c67b
	"api/room/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/kick", handler.KickUser)
	r.POST("/mute", handler.MuteUser)
	r.POST("/unmute", handler.UnmuteUser)

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

			// 麦位管理相关路由
			room.POST("/applyMic", handler.ApplyMic)
			room.POST("/leaveMic", handler.LeaveMic)

		}
	}
}
