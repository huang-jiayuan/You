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

// KickUser 踢出用户
func KickUser(c *gin.Context) {
	var req request.KickUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": err.Error(),
		})
		return
	}

	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc服务连接失败",
			"data": err.Error(),
		})
		return
	}
	defer conn.Close()

	client := __.NewRoomClient(conn)
	resp, err := client.KickUser(c, &__.KickUserReq{
		RoomId:     req.RoomID,
		UserId:     req.UserID,
		OperatorId: req.OperatorID,
		Reason:     req.Reason,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "踢出用户失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "踢出用户成功",
		"data": resp,
	})
	return

}

// MuteUser 禁言接口
func MuteUser(c *gin.Context) {
	var req request.MuteUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": err.Error(),
		})
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc服务连接失败",
			"data": err.Error(),
		})
		return
	}
	defer conn.Close()
	client := __.NewRoomClient(conn)
	resp, err := client.MuteUser(c, &__.MuteUserReq{
		RoomId:       req.RoomID,
		UserId:       req.UserID,
		OperatorId:   req.OperatorID,
		DurationType: req.DurationType,
		Reason:       req.Reason,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "禁言用户失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "禁言用户成功",
		"data": resp,
	})
	return
}

// UnmuteUser 解除禁言接口
func UnmuteUser(c *gin.Context) {
	var req request.UnmuteUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证失败",
			"data": err.Error(),
		})
		return
	}
	//调用grpc
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc服务连接失败",
			"data": err.Error(),
		})
		return
	}
	defer conn.Close()

	client := __.NewRoomClient(conn)
	resp, err := client.UnmuteUser(c, &__.UnmuteUserReq{
		RoomId:     req.RoomID,
		UserId:     req.UserID,
		OperatorId: req.OperatorID,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "解除禁言失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "解除禁言成功",
		"data": resp,
	})
	return
}
