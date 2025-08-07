package handler

import (
	"api/pkg"
	"api/request"
	__ "api/user/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"path/filepath"
)

// 示例
func Stream(c *gin.Context) {
	var req request.Stream
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	sendsms, err := c1.Greet(c, &__.StreamReq{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "短信发送失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "短信发送成功",
		"data": sendsms,
	})
	return
}
func Sendsms(c *gin.Context) {
	var req request.SendSms
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	sendsms, err := c1.SendSms(c, &__.SendSmsRequest{
		Username: req.Username,
		Source:   req.Source,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "短信发送失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "短信发送成功",
		"data": sendsms,
	})
	return
}
func Login(c *gin.Context) {
	var req request.UserLogin
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	lastLoginIp := c.ClientIP()
	login, err := c1.UserLogin(c, &__.UserLoginRequest{
		Username:    req.Username,
		Code:        req.Code,
		LastLoginIp: lastLoginIp,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	claims := pkg.CustomClaims{ID: uint(login.Greet)}
	token, err := pkg.NewJWT("2211a").CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "服务器内部错误",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": map[string]interface{}{"token": token},
	})
}

func UserPassword(c *gin.Context) {
	var req request.UserPassword
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	login, err := c1.UserPassword(c, &__.UserPasswordRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	if login.Token != "" {
		c.SetCookie(
			"remember_token",
			login.Token,
			30*24*3600, // 30 天有效期
			"/", "", false, true,
		)
	}
	claims := pkg.CustomClaims{ID: uint(login.Greet)}
	token, err := pkg.NewJWT("2211a").CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": map[string]interface{}{"token": token},
	})
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("remember_token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "未授权"})
			return
		}
		conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c1 := __.NewUserClient(conn)
		// 调用 gRPC 验证令牌
		resp, err := c1.VerifyToken(c, &__.VerifyTokenRequest{
			Token: token,
		})
		if err != nil || !resp.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "令牌无效或已过期"})
			return
		}

		// 令牌有效，挂载用户名到上下文
		c.Set("username", resp.Username)
		c.Next()
	}
}
func UpdatePassword(c *gin.Context) {
	var req request.UpdatePassword
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	userId := c.GetUint("userId")
	fmt.Println(userId)
	update, err := c1.UpdatePassword(c, &__.UpdatePasswordRequest{
		UserId:   int64(userId),
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": update,
	})
}

func ImproveUserMessage(c *gin.Context) {
	var req request.ImproveUserMessage
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewUserClient(conn)
	userId := c.GetUint("userId")

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "请求失败！无法找到获取的资源",
			"data": err.Error(),
		})
		return
	}
	fileHeader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileHeader.Close()
	ext := filepath.Ext(file.Filename)
	if ext != ".mp4" && ext != ".jpg" && ext != ".png" && ext != ".webp" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": "文件格式错误",
		})
		return
	}

	if file.Size >= 500*1024*1024 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": "文件过大",
		})
		return
	}

	upload, err := pkg.Upload(file.Filename, fileHeader, file.Size)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": "文件上传失败",
		})
		return
	}

	update, err := c1.ImproveUserMessage(c, &__.ImproveUserMessageRequest{
		UserId:       int64(userId),
		Nickname:     req.Nickname,
		Avatar:       upload,
		Gender:       req.Gender,
		VoiceTag:     req.VoiceTag,
		InterestTags: req.InterestTag,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "请求失败！服务器内部错误",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "服务器响应正常",
		"data": update,
	})
}

func Profile(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(200, gin.H{
		"username": username,
		"message":  "已通过认证",
	})
}

func Logout(c *gin.Context) {
	token, err := c.Cookie("remember_token")
	if err == nil {
		// 调用 gRPC 登出
		conn, err := grpc.NewClient("127.0.0.1:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c1 := __.NewUserClient(conn)
		_, _ = c1.Logout(c, &__.LogoutRequest{
			Token: token,
		})
		// 清除 Cookie
		c.SetCookie("remember_token", "", -1, "/", "", false, true)
	}

	c.JSON(200, gin.H{"success": true, "message": "登出成功"})
}
