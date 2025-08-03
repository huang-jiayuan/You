package handler

import (
	"api/pkg"
	"api/request"
	__ "api/user/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
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
		Mobile: req.Mobile,
		Source: req.Source,
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
	login, err := c1.UserLogin(c, &__.UserLoginRequest{
		Mobile: req.Mobile,
		Code:   req.Code,
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
