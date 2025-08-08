package router

import (
	"api/room/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/kick", handler.KickUser)
	r.POST("/mute", handler.MuteUser)
	r.POST("/unmute", handler.UnmuteUser)

}
