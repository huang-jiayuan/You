package main

import (
	"api/pkg"
	"api/user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pkg.InitMinio()

	// 添加CORS中间件 - 支持前端开发环境（包含IPv6支持）
	r.Use(pkg.CORSWithConfig([]string{
		"http://localhost:3000",   // 前端开发服务器
		"http://127.0.0.1:3000",  // IPv4本地回环地址
		"http://[::1]:3000",      // IPv6本地回环地址
		"*", // 开发环境允许所有来源，生产环境应该移除
	}))

	router.Router(r)
	// 强制绑定到IPv4地址，避免IPv6问题
	r.Run("127.0.0.1:8081")
}
