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
			room.POST("/createroom", handler.CreateRoom)
			room.POST("/updateroom", handler.UpdateRoom)
			room.POST("/getRecommendRooms", handler.GetRecommendRooms)
			room.POST("/getRoomsByCategory", handler.GetRoomsByCategory)
			room.POST("/searchRooms", handler.SearchRooms)

		}
	}
}
