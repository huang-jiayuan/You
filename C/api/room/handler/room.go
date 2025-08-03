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
func CreateRoom(c *gin.Context) {
	var req request.CreateRoom
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
	create, err := c1.CreateRoom(c, &__.CreateRoomStreamReq{
		RoomName: req.RoomName,
		UserId:   int32(req.UserId),
		TagId:    int32(req.TagId),
		Content:  req.Content,
		IdCard:   req.IdCard,
		RealName: req.RealName,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "房间创建失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "房间创建成功",
		"data": create,
	})
	return
}

func UpdateRoom(c *gin.Context) {
	var req request.UpdateRoom
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
	create, err := c1.UpdateRoom(c, &__.UpdateRoomStreamReq{
		Announcement: req.Announcement,
		Id:           int32(req.Id),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "房间信息修改失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "房间信息修改成功",
		"data": create,
	})
	return
}

func GetRecommendRooms(c *gin.Context) {
	var req request.GetRecommendRooms
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
	create, err := c1.GetRecommendRooms(c, &__.GetRecommendRoomsReq{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "房间展示失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "房间展示成功",
		"data": create,
	})
	return
}

func GetRoomsByCategory(c *gin.Context) {
	var req request.GetRoomsByCategory
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
	create, err := c1.GetRoomsByCategory(c, &__.GetRoomsByCategoryReq{
		TagId:    uint64(req.TagId),
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "房间展示失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "房间展示成功",
		"data": create,
	})
	return
}

func SearchRooms(c *gin.Context) {
	var req request.SearchRooms
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
	create, err := c1.SearchRooms(c, &__.SearchRoomsReq{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "未查找到有关房间",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "房间搜索成功",
		"data": create,
	})
	return
}
