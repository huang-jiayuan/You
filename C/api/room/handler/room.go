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

func SendGifts(c *gin.Context) {
	var req request.SendGifts
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "未接收参数",
			"data": nil,
		})
		return
	}
	client, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer client.Close()
	c1 := __.NewRoomClient(client)
	gifts, err := c1.SendGifts(c, &__.SendGiftsReq{
		SendUserId:    req.SendUserId,
		ReceiveUserId: req.ReceiveUserId,
		RoomId:        req.RoomId,
		GiftId:        req.GiftId,
		SendCount:     req.SendCount,
		SendType:      req.SendType,
		Message:       req.Message,
		Status:        req.Status,
		ClientIp:      req.ClientIp,
		SendTime:      req.SendTime,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "刷礼记录失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "刷礼记录成",
		"data": gifts,
	})
}
