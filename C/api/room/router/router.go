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
			room.Use(pkg.JWTAuth("2211a"))
			room.POST("/createroom", handler.CreateRoom)
			room.POST("/updateroom", handler.UpdateRoom)
			room.POST("/getRecommendRooms", handler.GetRecommendRooms)
			room.POST("/getRoomsByCategory", handler.GetRoomsByCategory)
			room.POST("/searchRooms", handler.SearchRooms)
			room.POST("/joinRoom", handler.JoinRoom)
			room.POST("/closeRoom", handler.CloseRoom)
			room.POST("/applyMic", handler.ApplyMic)
			room.POST("/leaveMic", handler.LeaveMic)

			room.POST("/handleMicApplication", handler.HandleMicApplication)
			room.POST("/kickFromMic", handler.KickFromMic)
			room.POST("/muteMicUser", handler.MuteMicUser)
		}
	}
}
