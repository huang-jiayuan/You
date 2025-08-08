package handler

import (
	"api/request"
	__ "api/room/proto"
	"fmt"
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

	// 添加调试信息
	userId := c.GetUint("userId")
	fmt.Printf("JWT解析的用户ID: %d, 创建房间: %s\n", userId, req.RoomName)

	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewRoomClient(conn)
	create, err := c1.CreateRoom(c, &__.CreateRoomStreamReq{
		RoomName: req.RoomName,
		UserId:   int32(c.GetUint("userId")), // 修改为userId
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
		fmt.Printf("SearchRooms参数绑定失败: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "参数验证失败",
			"data": err.Error(),
		})
		return
	}

	if len(req.Keyword) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "搜索关键词不能为空",
			"data": nil,
		})
		return
	}

	if len(req.Keyword) > 50 {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "搜索关键词长度不能超过50个字符",
			"data": nil,
		})
		return
	}

	fmt.Printf("搜索房间请求 - 关键词: %s, 页码: %d, 页大小: %d\n", req.Keyword, req.Page, req.PageSize)

	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewRoomClient(conn)
	create, err := c1.SearchRooms(c, &__.SearchRoomsReq{
		Keyword:  req.Keyword,
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "房间搜索失败",
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

func CloseRoom(c *gin.Context) {
	var req request.CloseRoom
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}

	// 添加调试信息
	userId := c.GetUint("userId")
	fmt.Printf("JWT解析的用户ID: %d, 请求关闭房间ID: %d\n", userId, req.RoomId)

	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c1 := __.NewRoomClient(conn)
	create, err := c1.CloseRoom(c, &__.CloseRoomStreamReq{
		RoomId: req.RoomId,
		UserId: uint64(c.GetUint("userId")), // 修改为userId
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "关闭房间失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "房间关闭成功",
		"data": create,
	})
	return
}

func JoinRoom(c *gin.Context) {
	var req request.JoinRoom
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
	create, err := c1.JoinRoom(c, &__.JoinRoomStreamReq{
		RoomId: req.RoomId,
		UserId: uint64(c.GetUint("userId")), // 修改为userId
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "进入房间失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "进入房间成功",
		"data": create,
	})
	return
}

// 申请上麦
func ApplyMic(c *gin.Context) {
	var req request.ApplyMic
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}

	// 获取JWT中的用户ID
	userId := c.GetUint("userId")
	fmt.Printf("JWT解析的用户ID: %d, 申请房间 %d 的麦位\n", userId, req.RoomId)

	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c1 := __.NewRoomClient(conn)
	result, err := c1.ApplyMic(c, &__.ApplyMicReq{
		RoomId: req.RoomId,
		UserId: uint64(userId),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "申请上麦失败",
			"data": err.Error(),
		})
		return
	}

	if !result.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  result.Message,
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  result.Message,
		"data": gin.H{
			"application_id": result.ApplicationId,
		},
	})
	return
}

// 下麦
func LeaveMic(c *gin.Context) {
	var req request.LeaveMic
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}

	// 获取JWT中的用户ID
	userId := c.GetUint("userId")
	fmt.Printf("JWT解析的用户ID: %d, 请求下麦房间 %d\n", userId, req.RoomId)

	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c1 := __.NewRoomClient(conn)
	result, err := c1.LeaveMic(c, &__.LeaveMicReq{
		RoomId: req.RoomId,
		UserId: uint64(userId),
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "下麦失败",
			"data": err.Error(),
		})
		return
	}

	if !result.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  result.Message,
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  result.Message,
		"data": gin.H{
			"mic_position": result.MicPosition,
		},
	})
	return
}

func HandleMicApplication(c *gin.Context) {
	var req request.HandleMicApplication
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
	create, err := c1.HandleMicApplication(c, &__.HandleMicApplicationReq{
		ApplicationId: req.ApplicationId,
		HandlerId:     req.HandlerId,
		Action:        req.Action,
		Reason:        req.Reason,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "处理申请失败",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "处理申请成功",
		"data": create,
	})
	return
}

func KickFromMic(c *gin.Context) {
	var req request.KickFromMic
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	
	// 获取JWT中的用户ID作为操作者ID
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "用户未登录或token无效",
			"data": nil,
		})
		return
	}
	
	// 验证被踢用户ID不能为0
	if req.TargetUserId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "被踢用户ID不能为空",
			"data": nil,
		})
		return
	}
	
	// 不能踢自己
	if uint64(userId) == req.TargetUserId {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "不能踢自己下麦",
			"data": nil,
		})
		return
	}
	
	fmt.Printf("踢人下麦 - 操作者ID: %d, 房间ID: %d, 被踢用户ID: %d, 原因: %s\n", 
		userId, req.RoomId, req.TargetUserId, req.Reason)
	
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	c1 := __.NewRoomClient(conn)
	result, err := c1.KickFromMic(c, &__.KickFromMicReq{
		RoomId:       req.RoomId,
		OperatorId:   uint64(userId),
		TargetUserId: req.TargetUserId,
		Reason:       req.Reason,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "踢人下麦失败",
			"data": err.Error(),
		})
		return
	}
	
	if !result.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  result.Message,
			"data": nil,
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  result.Message,
		"data": gin.H{
			"mic_position": result.MicPosition,
		},
	})
	return
}

func MuteMicUser(c *gin.Context) {
	var req request.MuteMicUser
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "验证失败",
			"data": err.Error(),
		})
		return
	}
	
	// 获取JWT中的用户ID作为操作者ID
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "用户未登录或token无效",
			"data": nil,
		})
		return
	}
	
	// 验证被操作用户ID不能为0
	if req.TargetUserId == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "目标用户ID不能为空",
			"data": nil,
		})
		return
	}
	
	// 不能操作自己
	if uint64(userId) == req.TargetUserId {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "不能对自己进行禁言操作",
			"data": nil,
		})
		return
	}
	
	fmt.Printf("禁言管理 - 操作者ID: %d, 房间ID: %d, 目标用户ID: %d, 动作: %d, 时长: %d分钟, 原因: %s\n", 
		userId, req.RoomId, req.TargetUserId, req.Action, req.Duration, req.Reason)
	
	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	c1 := __.NewRoomClient(conn)
	result, err := c1.MuteMicUser(c, &__.MuteMicUserReq{
		RoomId:       req.RoomId,
		OperatorId:   uint64(userId),
		TargetUserId: req.TargetUserId,
		Action:       req.Action,
		Duration:     req.Duration,
		Reason:       req.Reason,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "禁言/解禁失败",
			"data": err.Error(),
		})
		return
	}
	
	if !result.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  result.Message,
			"data": nil,
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  result.Message,
		"data": nil,
	})
	return
}
