package main

import (
	"api/user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run(":8081")
}
