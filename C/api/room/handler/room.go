package handler

import (
	"api/request"
	__ "api/room/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

// 示例
func Stream(c *gin.Context) {
	var req request.Streams
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewRoomClient(conn)
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
