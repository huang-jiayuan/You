package main

import (
	"api/room/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
}
