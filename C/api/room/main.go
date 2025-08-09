package main

import (
	"api/pkg"
	"api/room/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 添加CORS中间件
	r.Use(pkg.CORSMiddleware())

	router.Router(r)
	r.Run(":8080")
}
