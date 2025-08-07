package main

import (
	"api/pkg"
	"api/user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	// 添加CORS中间件
	r.Use(pkg.CORSMiddleware())
	
	router.Router(r)
	r.Run(":8081")
}
