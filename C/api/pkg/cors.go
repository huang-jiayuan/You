package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有来源，生产环境建议指定具体域名
		c.Header("Access-Control-Allow-Origin", "*")
		
		// 允许携带认证信息
		c.Header("Access-Control-Allow-Credentials", "true")
		
		// 允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		
		// 允许的请求方法
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		
		// 预检请求的缓存时间
		c.Header("Access-Control-Max-Age", "86400")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// CORSWithConfig 带配置的跨域中间件
func CORSWithConfig(allowOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// 检查是否在允许的来源列表中
		allowed := false
		for _, allowOrigin := range allowOrigins {
			if allowOrigin == "*" || allowOrigin == origin {
				allowed = true
				break
			}
		}
		
		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}