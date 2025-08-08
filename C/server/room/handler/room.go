package handler

import (
	"context"
<<<<<<< HEAD
	"server/models"
	__ "server/room/proto"
	"server/user/basic/global"
=======
	"fmt"
	"gorm.io/gorm"
	"server/models"
	"server/pkg"
	"server/room/basic/global"
	__ "server/room/proto"
>>>>>>> c0418eef7b50d6eb94c6431ae71816618e38c67b
	"time"
)

type Server struct {
	__.UnimplementedRoomServer
}

// KickUser 踢出用户接口
func (s *Server) KickUser(_ context.Context, req *__.KickUserReq) (*__.KickUserResp, error) {
	// 验证参数不能为空
	if req.RoomId == 0 || req.UserId == 0 || req.OperatorId == 0 {
		return &__.KickUserResp{
			Success: false,
			Message: "参数不能为空",
		}, nil
	}

	// 判断房主和踢出用户不是同一人
	if req.OperatorId == req.UserId {
		return &__.KickUserResp{
			Success: false,
			Message: "不能踢出自己",
		}, nil
	}

	// 判断用户是否已经被踢且在冷却时间内
	if models.IsUserKicked(req.RoomId, req.UserId) {
		return &__.KickUserResp{
			Success: false,
			Message: "用户已被踢出，在冷却时间内",
		}, nil
	}

	// 创建踢人记录
	now := time.Now()
	kickRecord := models.RoomKick{
		RoomID:     req.RoomId,
		UserID:     req.UserId,
		OperatorID: req.OperatorId,
		Reason:     req.Reason,
		KickTime:   now,
		ExpireTime: now.Add(10 * time.Minute), // 10分钟时间
	}

	// 创建踢出用户记录到数据库
	if err := global.DB.Create(&kickRecord).Error; err != nil {
		return &__.KickUserResp{
			Success: false,
			Message: "创建踢人记录失败: " + err.Error(),
		}, nil
	}

	return &__.KickUserResp{
		Success: true,
		Message: "用户已踢出房间，10分钟内不能再次进入",
	}, nil
}

// MuteUser 禁言接口
func (s *Server) MuteUser(_ context.Context, req *__.MuteUserReq) (*__.MuteUserResp, error) {
	// 验证参数是否为空
	if req.RoomId == 0 || req.UserId == 0 || req.OperatorId == 0 {
		return &__.MuteUserResp{
			Success: false,
			Message: "参数不能为空",
		}, nil
	}

	// 检查房主和被踢用户不是同一人
	if req.OperatorId == req.UserId {
		return &__.MuteUserResp{
			Success: false,
			Message: "不能禁言自己",
		}, nil
	}

	// 验证禁言时长类型
	if req.DurationType < 1 || req.DurationType > 3 {
		return &__.MuteUserResp{
			Success: false,
			Message: "禁言时长类型无效(1小时，2-24小时，3-永久)",
		}, nil
	}

	// 创建禁言记录
	now := time.Now()
	muteRecord := models.RoomMute{
		RoomID:       req.RoomId,
		UserID:       req.UserId,
		OperatorID:   req.OperatorId,
		DurationType: req.DurationType,
		Reason:       req.Reason,
		MuteTime:     now,
		ExpireTime:   models.GetMuteDuration(req.DurationType),
		IsActive:     true,
	}

	if err := global.DB.Create(&muteRecord).Error; err != nil {
		return &__.MuteUserResp{
			Success: false,
			Message: "创建禁言记录失败: " + err.Error(),
		}, nil
	}

	var durationMsg string
	switch req.DurationType {
	case 1:
		durationMsg = "1小时"
	case 2:
		durationMsg = "24小时"
	case 3:
		durationMsg = "永久"
	}

	return &__.MuteUserResp{
		Success: true,
		Message: "用户已被禁言" + durationMsg,
	}, nil
}

// UnmuteUser 解除禁言接口
func (s *Server) UnmuteUser(_ context.Context, req *__.UnmuteUserReq) (*__.UnmuteUserResp, error) {
	// 验证参数不能为空
	if req.RoomId == 0 || req.UserId == 0 || req.OperatorId == 0 {
		return &__.UnmuteUserResp{
			Success: false,
			Message: "参数不能为空",
		}, nil
	}

	// 判断房主和被踢用户不是同一人
	if req.OperatorId == req.UserId {
		return &__.UnmuteUserResp{
			Success: false,
			Message: "不能对自己禁言",
		}, nil
	}

	// 判断用户是否被禁言
	if !models.IsUserMuted(req.RoomId, req.UserId) {
		return &__.UnmuteUserResp{
			Success: false,
			Message: "用户未被禁言",
		}, nil
	}

	// 取消禁言记录
	if err := global.DB.Model(&models.RoomMute{}).Where("room_id = ? and user_id = ? and is_active = ?",
		req.RoomId, req.UserId, true).Update("is_active", false).Error; err != nil {
		return &__.UnmuteUserResp{
			Success: false,
			Message: "解除禁言失败: " + err.Error(),
		}, nil
	}

	return &__.UnmuteUserResp{
		Success: true,
		Message: "用户禁言已解除",
	}, nil
}

func (s *Server) CreateRoom(_ context.Context, in *__.CreateRoomStreamReq) (*__.CreateRoomStreamResp, error) {
	// 查询用户存不存在
	var user models.User
	err := global.DB.Where("id = ?", in.UserId).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	// 检查用户状态
	if user.Status != "0" {
		return nil, fmt.Errorf("账号状态异常，无法创建房间")
	}

	// 如果在没有实名的状态下
	if user.IsAuth == "0" {
		// 需要进行实名认证
		// 调用实名认证接口
		isValid, age, err := pkg.RealName(in.IdCard, in.RealName)
		if err != nil {
			return nil, fmt.Errorf("实名认证失败: %v", err)
		}
		if !isValid {
			return nil, fmt.Errorf("实名认证失败，请检查身份证号和姓名是否正确")
		}

		//年龄限制
		if age < 18 {
			return nil, fmt.Errorf("未满18岁，无法创建房间")
		}
		// 通过之后更新用户实名认证状态为1
		err = global.DB.Model(&user).Update("is_auth", 1).Error
		if err != nil {
			return nil, fmt.Errorf("更新实名认证状态失败")
		}
	}

	// 检查用户是否已经有活跃房间
	var count int64
	// 统计用户活跃房间数量（排除软删除的房间）
	err = global.DB.Model(&models.Room{}).Where("user_id = ? AND (closed_at IS NULL) AND status != '1' AND deleted_at IS NULL", in.UserId).Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("查询用户房间信息失败: %v", err)
	}
	// 如果大于0 说明有活跃房间
	if count > 0 {
		var activeRoom models.Room
		//查询如果关闭时间为空并且房间状态不是已关闭且未被软删除 说明有活跃房间
		err = global.DB.Where("user_id = ? AND closed_at IS NULL AND status != '1' AND deleted_at IS NULL", in.UserId).First(&activeRoom).Error
		if err != nil {
			return nil, fmt.Errorf("查询活跃房间失败")
		}
		return nil, fmt.Errorf("您已经有一个活跃的房间了（房间名称: %s），请先关闭现有房间后再创建新房间", activeRoom.RoomName)
	}
	//创建房间
	add := models.Room{
		RoomName: in.RoomName,
		UserId:   uint64(in.UserId),
		TagId:    uint64(in.TagId),
	}
	//只有不为空才检测敏感词
	if in.RoomName != "" {
		//调用敏感词检测
		isClean := pkg.Textcen(in.RoomName)
		if !isClean {
			return nil, fmt.Errorf("房间名称包含敏感信息，请修改后重试")
		}
	} else {
		fmt.Printf("房间名称为空，跳过敏感词检测\n")
	}
	err = global.DB.Create(&add).Error
	if err != nil {
		return nil, fmt.Errorf("房间创建失败: %v", err)
	}

	return &__.CreateRoomStreamResp{
		Id: int32(add.Id),
	}, nil
}
func (s *Server) UpdateRoom(_ context.Context, in *__.UpdateRoomStreamReq) (*__.UpdateRoomStreamResp, error) {
	var room models.Room
	err := global.DB.Where("id = ?", in.Id).First(&room).Error
	if err != nil {
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}
	// 验证房间状态
	if room.Status != "0" {
		return nil, fmt.Errorf("房间状态异常，无法修改公告")
	}

	// 验证房间是否已关闭
	if !room.ClosedAt.IsZero() {
		return nil, fmt.Errorf("房间已关闭，无法修改公告")
	}

	// 验证房间是否已被软删除
	if !room.DeletedAt.IsZero() {
		return nil, fmt.Errorf("房间已被删除，无法修改公告")
	}

	// 公告长度验证
	if len(in.Announcement) > 200 {
		return nil, fmt.Errorf("公告长度不能超过200个字符")
	}

	// 敏感词检测（仅当公告不为空时）
	if in.Announcement != "" {
		isClean := pkg.Textcen(in.Announcement)
		if !isClean {
			return nil, fmt.Errorf("房间公告包含敏感信息，请修改后重试")
		}
	} else {
		fmt.Printf("房间公告为空，跳过敏感词检测\n")
	}

	// 使用事务确保数据一致性
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新房间公告和相关字段
	err = tx.Model(&room).Updates(map[string]interface{}{
		"announcement":     in.Announcement,
		"updated_at":       time.Now(),
		"last_active_time": time.Now(),
	}).Error

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新房间公告失败: %v", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("保存房间公告失败: %v", err)
	}

	return &__.UpdateRoomStreamResp{
		Id: int32(room.Id),
	}, nil
}

// 分页
func Paginate(page, pageSize int32) (func(db *gorm.DB) *gorm.DB, int32, int32) {
	// 参数验证和默认值设置
	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize > 50:
		pageSize = 50
	case pageSize <= 0:
		pageSize = 10
	}

	scopeFunc := func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}

	return scopeFunc, page, pageSize
}

// 首页推荐
func (s *Server) GetRecommendRooms(ctx context.Context, req *__.GetRecommendRoomsReq) (*__.GetRecommendRoomsResp, error) {
	// 查询总数
	var total int64
	//房间状态正常并且关闭时间为空且未被软删除
	x := "status = '0' AND closed_at IS NULL AND deleted_at IS NULL"
	err := global.DB.Model(&models.Room{}).Where(x).Count(&total).Error
	if err != nil {
		return nil, fmt.Errorf("查询房间总数失败: %v", err)
	}

	// 查询房间列表
	paginateScope, page, pageSize := Paginate(req.Page, req.PageSize)
	var rooms []*models.Room
	err = global.DB.Scopes(paginateScope).Where(x).Order("fk_member_room DESC").Find(&rooms).Error
	if err != nil {
		return nil, fmt.Errorf("查询推荐房间失败: %v", err)
	}

	// 转换为RoomInfo
	var roomInfos []*__.RoomInfo
	for _, room := range rooms {
		// 查询房主信息
		var user models.User
		global.DB.Where("id = ?", room.UserId).First(&user)
		// 查询标签信息
		var tag models.RoomTagDict
		global.DB.Where("id = ?", room.TagId).First(&tag)

		roomInfo := &__.RoomInfo{
			Id:            room.Id,
			RoomName:      room.RoomName,
			UserId:        room.UserId,
			OwnerNickname: user.Nickname,
			OwnerAvatar:   user.Avatar,
			RoomType:      room.RoomType,
			TagId:         room.TagId,
			TagName:       tag.TagName, // 直接使用查询到的标签名称
		}
		roomInfos = append(roomInfos, roomInfo)
	}
	return &__.GetRecommendRoomsResp{
		Rooms:    roomInfos,
		Total:    int32(total),
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// 按房间标签分类
func (s *Server) GetRoomsByCategory(ctx context.Context, req *__.GetRoomsByCategoryReq) (*__.GetRoomsByCategoryResp, error) {

	// 先查询分类名称在不在
	var tag models.RoomTagDict
	err := global.DB.Where("id = ?", req.TagId).Find(&tag).Error
	if err != nil {
		return nil, fmt.Errorf("分类不存在")
	}
	// 查询总数
	var total int64
	//根据标签id查 并且房间状态正常且未被软删除
	c := "status = '0' AND closed_at IS NULL AND deleted_at IS NULL AND tag_id = ?"
	err = global.DB.Model(&models.Room{}).Where(c, req.TagId).Count(&total).Error
	if err != nil {
		return nil, fmt.Errorf("查询房间总数失败: %v", err)
	}
	// 查询房间列表
	paginateScope, page, pageSize := Paginate(req.Page, req.PageSize)
	var rooms []*models.Room
	err = global.DB.Scopes(paginateScope).Where(c, req.TagId).Order("fk_member_room DESC").Find(&rooms).Error
	if err != nil {
		return nil, fmt.Errorf("查询分类房间失败: %v", err)
	}
	var roomInfos []*__.RoomInfo
	for _, room := range rooms {
		// 查询房主信息
		var user models.User
		global.DB.Where("id = ?", room.UserId).Find(&user)
		roomInfo := &__.RoomInfo{
			Id:            room.Id,
			RoomName:      room.RoomName,
			UserId:        room.UserId,
			OwnerNickname: user.Nickname,
			OwnerAvatar:   user.Avatar,
			RoomType:      room.RoomType,
			TagId:         room.TagId,
			TagName:       tag.TagName, // 直接使用查询到的标签名称
		}
		roomInfos = append(roomInfos, roomInfo)
	}

	return &__.GetRoomsByCategoryResp{
		Rooms:        roomInfos,
		Total:        int32(total),
		Page:         page,
		PageSize:     pageSize,
		CategoryName: tag.TagName,
	}, nil
}

// 支持房间名称和房主昵称搜索
func (s *Server) SearchRooms(ctx context.Context, req *__.SearchRoomsReq) (*__.SearchRoomsResp, error) {
	if len(req.Keyword) > 50 {
		return nil, fmt.Errorf("搜索关键词长度不能超过50个字符")
	}

	// 模糊查询条件
	searchPattern := "%" + req.Keyword + "%"
	z := "status = '0' AND closed_at IS NULL AND deleted_at IS NULL AND room_name LIKE ?"

	// 查询总数
	var total int64
	err := global.DB.Model(&models.Room{}).Where(z, searchPattern).Count(&total).Error
	if err != nil {
		return nil, fmt.Errorf("查询搜索结果总数失败: %v", err)
	}

	// 查询房间列表
	paginateScope, page, pageSize := Paginate(req.Page, req.PageSize)
	var rooms []*models.Room
	//根据在线人数倒序展示
	err = global.DB.Scopes(paginateScope).Where(z, searchPattern).Order("fk_member_room DESC, created_at DESC").
		Find(&rooms).Error
	if err != nil {
		return nil, fmt.Errorf("搜索房间失败: %v", err)
	}

	fmt.Printf("实际查询到的房间数量: %d\n", len(rooms))

	// 转换为RoomInfo
	var roomInfos []*__.RoomInfo
	for _, room := range rooms {
		// 查询房主信息
		var user models.User
		global.DB.Where("id = ?", room.UserId).First(&user)
		// 查询标签信息
		var tag models.RoomTagDict
		global.DB.Where("id = ?", room.TagId).First(&tag)

		roomInfo := &__.RoomInfo{
			Id:            room.Id,
			RoomName:      room.RoomName,
			UserId:        room.UserId,
			OwnerNickname: user.Nickname,
			OwnerAvatar:   user.Avatar,
			RoomType:      room.RoomType,
			TagId:         room.TagId,
			TagName:       tag.TagName, // 直接使用查询到的标签名称
		}
		roomInfos = append(roomInfos, roomInfo)
	}
	return &__.SearchRoomsResp{
		Rooms:    roomInfos,
		Total:    int32(total),
		Page:     page,
		PageSize: pageSize,
		Keyword:  req.Keyword,
	}, nil
}

// 关闭房间接口
func (s *Server) CloseRoom(ctx context.Context, req *__.CloseRoomStreamReq) (*__.CloseRoomStreamResp, error) {

	// 验证房间是否存在
	var room models.Room
	err := global.DB.Where("id = ?", req.RoomId).Find(&room).Error
	if err != nil {
		return nil, fmt.Errorf("房间不存在 %v", err)
	}

	// 验证用户是否是房主
	if room.UserId != req.UserId {
		return nil, fmt.Errorf("只有房主才能关闭房间 %v", err)
	}

	// 检查房间是否已经关闭
	if room.Status == "1" || !room.ClosedAt.IsZero() {
		return nil, fmt.Errorf("房间已关闭 %v", err)
	}

	// 检查房间是否已经被软删除
	if !room.DeletedAt.IsZero() {
		return nil, fmt.Errorf("房间已被删除 %v", err)
	}

	//开始事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//将所有在线用户设为离线
	err = tx.Model(&models.UserRoom{}).Where("room_id = ? AND is_online = 1", req.RoomId).Updates(map[string]interface{}{
		"is_online":  0,
		"leave_time": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return &__.CloseRoomStreamResp{
			Success: false,
			Message: "更新用户状态失败",
			RoomId:  req.RoomId,
		}, nil
	}

	// 更新房间状态为已解散并进行软删除
	err = tx.Model(&room).Where("id = ?", req.RoomId).Updates(map[string]interface{}{
		"status":           "1",
		"closed_at":        time.Now(),
		"fk_member_room":   0, // 清零在线人数
		"last_active_time": time.Now(),
		"deleted_at":       time.Now(), // 软删除时间戳
		"deleted_by":       req.UserId, // 记录删除操作者
		"updated_at":       time.Now(), // 更新修改时间
		"updated_by":       req.UserId, // 记录修改者
	}).Error
	if err != nil {
		tx.Rollback()
		return &__.CloseRoomStreamResp{
			Success: false,
			Message: "更新房间状态失败",
			RoomId:  req.RoomId,
		}, nil
	}

	//  提交事务
	if err := tx.Commit().Error; err != nil {
		return &__.CloseRoomStreamResp{
			Success: false,
			Message: "关闭房间失败",
			RoomId:  req.RoomId,
		}, nil
	}

	return &__.CloseRoomStreamResp{
		Success: true,
		Message: "房间已成功关闭",
		RoomId:  req.RoomId,
	}, nil
}

// 进入房间接口
func (s *Server) JoinRoom(ctx context.Context, req *__.JoinRoomStreamReq) (*__.JoinRoomStreamResp, error) {
	var room models.Room
	//查询还存在并且正常且未被软删除的房间
	err := global.DB.Where("id = ? AND status = '0' AND closed_at IS NULL AND deleted_at IS NULL", req.RoomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("房间不存在或已关闭")
		}
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}
	// 检查用户是否已经在房间中
	var existingUserRoom models.UserRoom
	err = global.DB.Where("user_id = ? AND room_id = ? AND is_online = 1", req.UserId, req.RoomId).First(&existingUserRoom).Error
	if err == nil {
		// 用户已经在房间中
		fmt.Printf("用户 %d 已经在房间 %d 中\n", req.UserId, req.RoomId)

		// 查询当前房主和标签信息
		var roomOwner models.User
		global.DB.Where("id = ?", room.UserId).First(&roomOwner)

		var tag models.RoomTagDict
		global.DB.Where("id = ?", room.TagId).First(&tag)
		//返回当前房间信息
		roomInfo := &__.RoomInfo{
			Id:             room.Id,
			RoomName:       room.RoomName,
			UserId:         room.UserId,
			OwnerNickname:  roomOwner.Nickname,
			OwnerAvatar:    roomOwner.Avatar, //头像
			RoomType:       room.RoomType,
			Status:         room.Status,
			Announcement:   room.Announcement,
			CreatedAt:      room.CreatedAt.Format("2006-01-02 15:04:05"),
			LastActiveTime: room.LastActiveTime.Format("2006-01-02 15:04:05"),
			MemberCount:    room.FkMemberRoom,
			Image:          room.Image,
			TagId:          room.TagId,
			TagName:        tag.TagName,
		}

		return &__.JoinRoomStreamResp{
			Success:            true,
			Message:            "您已在房间中",
			CurrentMemberCount: room.FkMemberRoom,
			RoomInfo:           roomInfo,
		}, nil
	}

	// 开始事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 如果离开房间之后要修改状态
	z := "user_id = ? AND room_id = ?"
	err = tx.Model(&models.UserRoom{}).Where(z, req.UserId, req.RoomId).Updates(map[string]interface{}{
		"is_online":  0,
		"leave_time": time.Now(),
	}).Error
	//如果修改没有成功就回滚
	if err != nil {
		tx.Rollback()
		return &__.JoinRoomStreamResp{
			Success: false,
			Message: "更新用户状态失败",
		}, nil
	}

	// 如果再进入房间创建新的在线记录
	userRoom := models.UserRoom{
		UserId:   req.UserId,
		RoomId:   req.RoomId,
		IsOnline: 1,
		JoinTime: time.Now(),
	}
	err = tx.Create(&userRoom).Error
	if err != nil {
		tx.Rollback()
		return &__.JoinRoomStreamResp{
			Success: false,
			Message: "创建用户房间记录失败",
		}, nil
	}

	// 更新房间人数和最后活跃时间
	var onlineCount int64
	// 先统计当前在线人数
	err = tx.Model(&models.UserRoom{}).Where("room_id = ? AND is_online = 1", req.RoomId).Count(&onlineCount).Error
	if err != nil {
		tx.Rollback()
		return &__.JoinRoomStreamResp{
			Success: false,
			Message: "统计在线人数失败",
		}, nil
	}

	// 更新房间信息
	err = tx.Model(&room).Where("id = ?", req.RoomId).Updates(map[string]interface{}{
		"fk_member_room":   int32(onlineCount),
		"last_active_time": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return &__.JoinRoomStreamResp{
			Success: false,
			Message: "更新房间信息失败",
		}, nil
	}

	//提交事务
	if err := tx.Commit().Error; err != nil {
		return &__.JoinRoomStreamResp{
			Success: false,
			Message: "进入房间失败",
		}, nil
	}

	// 查询房主和标签信息
	var roomOwner models.User
	global.DB.Where("id = ?", room.UserId).First(&roomOwner)

	var tag models.RoomTagDict
	global.DB.Where("id = ?", room.TagId).First(&tag)

	roomInfo := &__.RoomInfo{
		Id:             room.Id,
		RoomName:       room.RoomName,
		UserId:         room.UserId,
		OwnerNickname:  roomOwner.Nickname,
		OwnerAvatar:    roomOwner.Avatar,
		RoomType:       room.RoomType,
		Status:         room.Status,
		Announcement:   room.Announcement,
		CreatedAt:      room.CreatedAt.Format("2006-01-02 15:04:05"),
		LastActiveTime: time.Now().Format("2006-01-02 15:04:05"), // 使用最新的活跃时间
		MemberCount:    int32(onlineCount),
		Image:          room.Image,
		TagId:          room.TagId,
		TagName:        tag.TagName,
	}

	fmt.Printf("用户 %d 成功进入房间 %d，当前在线人数: %d\n", req.UserId, req.RoomId, onlineCount)

	return &__.JoinRoomStreamResp{
		Success:            true,
		Message:            "成功进入房间",
		CurrentMemberCount: int32(onlineCount),
		RoomInfo:           roomInfo,
	}, nil
}

// 申请上麦功能
func (s *Server) ApplyMic(ctx context.Context, req *__.ApplyMicReq) (*__.ApplyMicResp, error) {
	// 验证房间存不存在
	var room models.Room
	err := global.DB.Where("id = ? AND status = '0' AND closed_at IS NULL AND deleted_at IS NULL", req.RoomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.ApplyMicResp{
				Success: false,
				Message: "房间不存在或已关闭",
			}, nil
		}
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}

	// 验证用户是否在房间中
	var userRoom models.UserRoom
	err = global.DB.Where("user_id = ? AND room_id = ? AND is_online = 1", req.UserId, req.RoomId).First(&userRoom).Error
	if err != nil {
		return &__.ApplyMicResp{
			Success: false,
			Message: "用户不在房间中",
		}, nil
	}

	// 先执行自动解禁过期用户
	s.autoUnmuteExpiredUsers()

	// 检查用户是否被禁言
	var roomMember models.RoomMember
	err = global.DB.Where("user_id = ? AND room_id = ?", req.UserId, req.RoomId).First(&roomMember).Error
	if err == nil {
		// 如果用户被禁言
		if roomMember.IsMuted == 1 {
			// 检查禁言是否已过期（双重检查，防止并发问题）
			if roomMember.MuteEndTime.IsZero() || time.Now().Before(roomMember.MuteEndTime) {
				var remainingTime time.Duration
				if !roomMember.MuteEndTime.IsZero() {
					remainingTime = time.Until(roomMember.MuteEndTime)
				}

				if remainingTime > 0 {
					return &__.ApplyMicResp{
						Success: false,
						Message: fmt.Sprintf("用户被禁言，无法申请上麦"),
					}, nil
				} else {
					return &__.ApplyMicResp{
						Success: false,
						Message: "用户被禁言，无法申请上麦",
					}, nil
				}
			}
			// 禁言已过期解除禁言
			global.DB.Model(&roomMember).Updates(map[string]interface{}{
				"is_muted":      0,
				"mute_end_time": nil,
			})
		}
	}

	// 检查用户是否已在麦位上
	var existingMic models.RoomMic
	err = global.DB.Where("room_id = ? AND user_id = ? AND status = 1", req.RoomId, req.UserId).First(&existingMic).Error
	if err == nil {
		return &__.ApplyMicResp{
			Success: false,
			Message: "用户已在麦位上",
		}, nil
	}

	//检查麦位是否已满
	var occupiedCount int64
	err = global.DB.Model(&models.RoomMic{}).Where("room_id = ? AND status = 1", req.RoomId).Count(&occupiedCount).Error
	if err != nil {
		return nil, fmt.Errorf("查询麦位状态失败: %v", err)
	}
	//计数
	if occupiedCount >= 8 {
		return &__.ApplyMicResp{
			Success: false,
			Message: "麦位已满，无法申请",
		}, nil
	}

	// 创建申请记录
	application := models.MicApplication{
		RoomId:    req.RoomId,
		UserId:    req.UserId,
		Status:    0, // 待审批
		ApplyTime: time.Now(),
	}

	err = global.DB.Create(&application).Error
	if err != nil {
		return nil, fmt.Errorf("创建申请记录失败: %v", err)
	}

	return &__.ApplyMicResp{
		Success:       true,
		Message:       "申请已提交，等待房主/管理员审批",
		ApplicationId: application.Id,
	}, nil
}

// 处理申请上麦功能
func (s *Server) HandleMicApplication(ctx context.Context, req *__.HandleMicApplicationReq) (*__.HandleMicApplicationResp, error) {
	// 查询申请记录并验证状态
	var application models.MicApplication
	err := global.DB.Where("id = ?", req.ApplicationId).First(&application).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.HandleMicApplicationResp{
				Success: false,
				Message: "申请记录不存在",
			}, nil
		}
		return nil, fmt.Errorf("查询申请记录失败: %v", err)
	}

	// 验证申请状态
	if application.Status != 0 {
		return &__.HandleMicApplicationResp{
			Success: false,
			Message: "申请已被处理",
		}, nil
	}

	// 验证房间存在性和状态
	var room models.Room
	err = global.DB.Where("id = ? AND status = '0' AND closed_at IS NULL AND deleted_at IS NULL", application.RoomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.HandleMicApplicationResp{
				Success: false,
				Message: "房间不存在或已关闭",
			}, nil
		}
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}

	// 首先处理人要存在房间里
	var roomMember models.RoomMember
	err = global.DB.Where("user_id = ? AND room_id = ?", req.HandlerId, application.RoomId).First(&roomMember).Error
	if err != nil {
		return &__.HandleMicApplicationResp{
			Success: false,
			Message: "处理者不在房间中",
		}, nil
	}
	// 检查权限 1=房主 2=管理员
	if roomMember.Role != 1 && roomMember.Role != 2 {
		return &__.HandleMicApplicationResp{
			Success: false,
			Message: "权限不足，只有房主或管理员可以处理申请",
		}, nil
	}

	// 另一种情况 用户申请完上麦之后离开房间
	var userRoom models.UserRoom
	err = global.DB.Where("user_id = ? AND room_id = ? AND is_online = 1", application.UserId, application.RoomId).First(&userRoom).Error
	if err != nil {
		// 用户已离开房间自动取消申请
		global.DB.Model(&application).Updates(map[string]interface{}{
			"status":      3, // 已取消
			"handle_time": time.Now(),
			"handler_id":  req.HandlerId,
			"reason":      "用户已离开房间",
		})
		return &__.HandleMicApplicationResp{
			Success: false,
			Message: "申请用户已离开房间，申请已自动取消",
		}, nil
	}

	// 开始事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if req.Action == 1 { // 批准申请
		// 检查用户是否已在麦位上
		var existingMic models.RoomMic
		err = tx.Where("room_id = ? AND user_id = ? AND status = 1", application.RoomId, application.UserId).First(&existingMic).Error
		if err == nil {
			tx.Rollback()
			return &__.HandleMicApplicationResp{
				Success: false,
				Message: "用户已在麦位上",
			}, nil
		}

		// 查找可用麦位
		var availableMic models.RoomMic
		//0是未上锁的麦
		err = tx.Where("room_id = ? AND status = 0 AND is_locked = 0", application.RoomId).Order("mic_position").First(&availableMic).Error
		if err != nil {
			tx.Rollback()
			return &__.HandleMicApplicationResp{
				Success: false,
				Message: "没有可用的麦位",
			}, nil
		}

		// 分配麦位
		err = tx.Model(&availableMic).Updates(map[string]interface{}{
			"status":      1, // 占用
			"user_id":     application.UserId,
			"occupy_time": time.Now(),
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("分配麦位失败: %v", err)
		}

		// 更新申请状态为已批准
		err = tx.Model(&application).Updates(map[string]interface{}{
			"status":      1, // 已批准
			"handle_time": time.Now(),
			"handler_id":  req.HandlerId,
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("更新申请状态失败: %v", err)
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			return nil, fmt.Errorf("处理申请失败: %v", err)
		}

		// 查询用户信息用于返回
		var user models.User
		global.DB.Where("id = ?", application.UserId).First(&user)

		// 返回成功结果和麦位信息
		micInfo := &__.MicInfo{
			Position:     int32(availableMic.MicPosition),
			Status:       1, // 占用
			UserId:       application.UserId,
			UserNickname: user.Nickname,
			UserAvatar:   user.Avatar,
			OccupyTime:   time.Now().Format("2006-01-02 15:04:05"),
			IsLocked:     availableMic.IsLocked == 1, //代表已锁
			IsMuted:      false,                      // 新上麦用户默认不禁言
		}

		return &__.HandleMicApplicationResp{
			Success: true,
			Message: "申请已批准，用户已上麦",
			MicInfo: micInfo,
		}, nil

	} else if req.Action == 2 { // 第二张情况就是拒绝申请
		// 更新申请状态为已拒绝
		err = tx.Model(&application).Updates(map[string]interface{}{
			"status":      2, // 已拒绝
			"handle_time": time.Now(),
			"handler_id":  req.HandlerId,
			"reason":      req.Reason,
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("更新申请状态失败: %v", err)
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			return nil, fmt.Errorf("处理申请失败: %v", err)
		}

		return &__.HandleMicApplicationResp{
			Success: true,
			Message: "申请已拒绝",
		}, nil

	} else {
		return &__.HandleMicApplicationResp{
			Success: false,
			Message: "无效的操作类型",
		}, nil
	}
}

// 下麦功能
func (s *Server) LeaveMic(ctx context.Context, req *__.LeaveMicReq) (*__.LeaveMicResp, error) {
	// 验证房间存不存在
	var room models.Room
	err := global.DB.Where("id = ? AND status = '0' AND closed_at IS NULL AND deleted_at IS NULL", req.RoomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.LeaveMicResp{
				Success: false,
				Message: "房间不存在或已关闭",
			}, nil
		}
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}

	// 查麦位管理表
	var roomMic models.RoomMic
	err = global.DB.Where("room_id = ? AND user_id = ? AND status = 1", req.RoomId, req.UserId).First(&roomMic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.LeaveMicResp{
				Success: false,
				Message: "用户不在麦位上",
			}, nil
		}
		return nil, fmt.Errorf("查询麦位状态失败: %v", err)
	}

	//开始事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 释放麦位并更新麦位管理表状态
	err = tx.Model(&roomMic).Updates(map[string]interface{}{
		"status":      0,   // 设置为空闲状态
		"user_id":     nil, // 清除用户ID
		"occupy_time": nil, // 清除占用时间信息
	}).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("释放麦位失败: %v", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("下麦操作失败: %v", err)
	}
	return &__.LeaveMicResp{
		Success:     true,
		Message:     "成功下麦",
		MicPosition: int32(roomMic.MicPosition),
	}, nil
}

// 踢人下麦功能 只有管理员和房主能操作
func (s *Server) KickFromMic(ctx context.Context, req *__.KickFromMicReq) (*__.KickFromMicResp, error) {
	// 验证房间存在
	var room models.Room
	err := global.DB.Where("id = ? AND status = '0' AND closed_at IS NULL AND deleted_at IS NULL", req.RoomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.KickFromMicResp{
				Success: false,
				Message: "房间不存在或已关闭",
			}, nil
		}
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}

	//先查操作人在不在房间里
	var operatorMember models.RoomMember
	err = global.DB.Where("user_id = ? AND room_id = ?", req.OperatorId, req.RoomId).First(&operatorMember).Error
	if err != nil {
		return &__.KickFromMicResp{
			Success: false,
			Message: "操作者不在房间中",
		}, nil
	}

	// 检查权限：房主(role=2)或管理员(role=1)
	if operatorMember.Role != 1 && operatorMember.Role != 2 {
		return &__.KickFromMicResp{
			Success: false,
			Message: "权限不足，只有房主或管理员可以踢人下麦",
		}, nil
	}

	// 验证目标用户当在不在前在麦位上
	var targetMic models.RoomMic
	err = global.DB.Where("room_id = ? AND user_id = ? AND status = 1", req.RoomId, req.TargetUserId).First(&targetMic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.KickFromMicResp{
				Success: false,
				Message: "目标用户不在麦位上",
			}, nil
		}
		return nil, fmt.Errorf("查询目标用户麦位状态失败: %v", err)
	}

	//开始事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//修改目标用户的麦位
	err = tx.Model(&targetMic).Updates(map[string]interface{}{
		"status":      0,   // 设置为空闲状态
		"user_id":     nil, // 清除用户ID
		"occupy_time": nil, // 清除占用时间信息
	}).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("释放麦位失败: %v", err)
	}

	// 记录踢人操作的详细日志
	// 日志解析和自动化处理 快速搜索特定操作或用户 生成统计报表 集成到日志分析系统中
	// 便于后续查询和审计特定的管理操作
	var operatorUser models.User
	var targetUser models.User
	tx.Where("id = ?", req.OperatorId).First(&operatorUser)
	tx.Where("id = ?", req.TargetUserId).First(&targetUser)

	// 记录操作日志到系统日志
	logMessage := fmt.Sprintf("踢人下麦操作 - 房间ID: %d, 操作者: %s(ID:%d), 目标用户: %s(ID:%d), 麦位: %d, 原因: %s, 时间: %s",
		req.RoomId,
		operatorUser.Nickname, req.OperatorId,
		targetUser.Nickname, req.TargetUserId,
		targetMic.MicPosition,
		req.Reason,
		time.Now().Format("2006-01-02 15:04:05"))

	fmt.Printf("[KICK_FROM_MIC] %s\n", logMessage)

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("踢人下麦操作失败: %v", err)
	}
	return &__.KickFromMicResp{
		Success:     true,
		Message:     fmt.Sprintf("成功将用户 %s 踢下麦位", targetUser.Nickname),
		MicPosition: int32(targetMic.MicPosition),
	}, nil
}

// 自动解除过期禁言的辅助函数
func (s *Server) autoUnmuteExpiredUsers() error {
	// 查询所有已过期的禁言用户
	var expiredMutedMembers []models.RoomMember
	//根据是否禁言 禁言结束时间
	err := global.DB.Where("is_muted = 1 AND mute_end_time IS NOT NULL AND mute_end_time <= ?", time.Now()).Find(&expiredMutedMembers).Error
	if err != nil {
		return fmt.Errorf("查询过期禁言用户失败: %v", err)
	}

	if len(expiredMutedMembers) == 0 {
		return nil // 没有过期的禁言用户
	}

	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 批量解除过期禁言
	for _, member := range expiredMutedMembers {
		err = tx.Model(&member).Updates(map[string]interface{}{
			"is_muted":      0,
			"mute_end_time": nil,
		}).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("自动解除用户 %d 禁言失败: %v", member.UserId, err)
		}
		// 记录自动解禁日志
		fmt.Printf("[AUTO_UNMUTE] 自动解除用户 %d 在房间 %d 的禁言，解禁时间: %s\n",
			member.UserId, member.RoomId, time.Now().Format("2006-01-02 15:04:05"))
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交自动解禁操作失败: %v", err)
	}

	fmt.Printf("[AUTO_UNMUTE] 成功自动解除 %d 个用户的过期禁言\n", len(expiredMutedMembers))
	return nil
}

// 禁言管理功能
func (s *Server) MuteMicUser(ctx context.Context, req *__.MuteMicUserReq) (*__.MuteMicUserResp, error) {
	// 先执行自动解禁过期用户
	s.autoUnmuteExpiredUsers()

	// 1. 验证房间存在性和状态
	var room models.Room
	err := global.DB.Where("id = ? AND status = '0' AND closed_at IS NULL AND deleted_at IS NULL", req.RoomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &__.MuteMicUserResp{
				Success: false,
				Message: "房间不存在或已关闭",
			}, nil
		}
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}

	// 2. 验证操作者权限（房主或管理员）
	var operatorMember models.RoomMember
	err = global.DB.Where("user_id = ? AND room_id = ?", req.OperatorId, req.RoomId).First(&operatorMember).Error
	if err != nil {
		return &__.MuteMicUserResp{
			Success: false,
			Message: "操作者不在房间中",
		}, nil
	}

	// 检查权限：房主(role=2)或管理员(role=1)
	if operatorMember.Role != 1 && operatorMember.Role != 2 {
		return &__.MuteMicUserResp{
			Success: false,
			Message: "权限不足，只有房主或管理员可以进行禁言操作",
		}, nil
	}

	// 3. 查询目标用户的房间成员记录
	var targetMember models.RoomMember
	err = global.DB.Where("user_id = ? AND room_id = ?", req.TargetUserId, req.RoomId).First(&targetMember).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果目标用户不在房间成员表中，创建一个记录
			targetMember = models.RoomMember{
				RoomId:   req.RoomId,
				UserId:   req.TargetUserId,
				Role:     0, // 普通成员
				JoinTime: time.Now(),
				IsMuted:  0,
			}
			err = global.DB.Create(&targetMember).Error
			if err != nil {
				return nil, fmt.Errorf("创建用户房间成员记录失败: %v", err)
			}
		} else {
			return nil, fmt.Errorf("查询目标用户房间成员记录失败: %v", err)
		}
	}

	// 4. 验证目标用户在房间中（检查在线状态）
	var userRoom models.UserRoom
	err = global.DB.Where("user_id = ? AND room_id = ? AND is_online = 1", req.TargetUserId, req.RoomId).First(&userRoom).Error
	if err != nil {
		return &__.MuteMicUserResp{
			Success: false,
			Message: "目标用户不在房间中",
		}, nil
	}

	// 5. 开始事务处理
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询操作者和目标用户信息用于日志记录
	var operatorUser models.User
	var targetUser models.User
	tx.Where("id = ?", req.OperatorId).First(&operatorUser)
	tx.Where("id = ?", req.TargetUserId).First(&targetUser)

	if req.Action == 1 { // 禁言操作
		// 6. 计算禁言结束时间
		var muteEndTime time.Time
		if req.Duration > 0 {
			muteEndTime = time.Now().Add(time.Duration(req.Duration) * time.Minute)
		} else {
			// 如果没有指定时长，默认禁言30分钟
			muteEndTime = time.Now().Add(30 * time.Minute)
		}

		// 7. 更新RoomMember表的禁言状态和结束时间
		err = tx.Model(&targetMember).Updates(map[string]interface{}{
			"is_muted":      1,
			"mute_end_time": muteEndTime,
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("更新用户禁言状态失败: %v", err)
		}

		// 8. 如果被禁言用户在麦位上，标记相应状态（但不踢下麦位，只是标记为禁言状态）
		var roomMic models.RoomMic
		err = tx.Where("room_id = ? AND user_id = ? AND status = 1", req.RoomId, req.TargetUserId).First(&roomMic).Error
		if err == nil {
			// 用户在麦位上，记录日志表明用户在麦位上被禁言
			fmt.Printf("[MUTE_MIC_USER] 用户 %s(ID:%d) 在麦位 %d 上被禁言\n",
				targetUser.Nickname, req.TargetUserId, roomMic.MicPosition)
		}

		// 9. 记录操作日志
		logMessage := fmt.Sprintf("禁言操作 - 房间ID: %d, 操作者: %s(ID:%d), 目标用户: %s(ID:%d), 禁言时长: %d分钟, 结束时间: %s, 原因: %s, 时间: %s",
			req.RoomId,
			operatorUser.Nickname, req.OperatorId,
			targetUser.Nickname, req.TargetUserId,
			req.Duration,
			muteEndTime.Format("2006-01-02 15:04:05"),
			req.Reason,
			time.Now().Format("2006-01-02 15:04:05"))

		fmt.Printf("[MUTE_MIC_USER] %s\n", logMessage)

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			return nil, fmt.Errorf("禁言操作失败: %v", err)
		}

		return &__.MuteMicUserResp{
			Success: true,
			Message: fmt.Sprintf("成功禁言用户 %s，禁言时长 %d 分钟", targetUser.Nickname, req.Duration),
		}, nil

	} else if req.Action == 2 { // 解禁操作
		// 10. 检查用户是否被禁言
		if targetMember.IsMuted != 1 {
			tx.Rollback()
			return &__.MuteMicUserResp{
				Success: false,
				Message: "用户当前未被禁言",
			}, nil
		}

		// 11. 更新RoomMember表，解除禁言状态
		err = tx.Model(&targetMember).Updates(map[string]interface{}{
			"is_muted":      0,
			"mute_end_time": nil,
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("解除用户禁言状态失败: %v", err)
		}

		// 12. 记录解禁操作日志
		logMessage := fmt.Sprintf("解禁操作 - 房间ID: %d, 操作者: %s(ID:%d), 目标用户: %s(ID:%d), 原因: %s, 时间: %s",
			req.RoomId,
			operatorUser.Nickname, req.OperatorId,
			targetUser.Nickname, req.TargetUserId,
			req.Reason,
			time.Now().Format("2006-01-02 15:04:05"))

		fmt.Printf("[UNMUTE_MIC_USER] %s\n", logMessage)

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			return nil, fmt.Errorf("解禁操作失败: %v", err)
		}

		return &__.MuteMicUserResp{
			Success: true,
			Message: fmt.Sprintf("成功解除用户 %s 的禁言", targetUser.Nickname),
		}, nil

	} else {
		tx.Rollback()
		return &__.MuteMicUserResp{
			Success: false,
			Message: "无效的操作类型，请使用 1-禁言 或 2-解禁",
		}, nil
	}
}

// GetMicStatus 获取麦位状态功能
// Author [yourname](https://github.com/yourname)
func (s *Server) GetMicStatus(ctx context.Context, req *__.GetMicStatusReq) (*__.GetMicStatusResp, error) {
	// 1. 验证房间存在性和状态
	var room models.Room
	err := global.DB.Where("id = ? AND status = '0' AND closed_at IS NULL AND deleted_at IS NULL", req.RoomId).First(&room).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("房间不存在或已关闭")
		}
		return nil, fmt.Errorf("查询房间失败: %v", err)
	}
	// 3. 查询房间所有麦位的当前状态
	var roomMics []models.RoomMic
	err = global.DB.Where("room_id = ?", req.RoomId).Order("mic_position").Find(&roomMics).Error
	if err != nil {
		return nil, fmt.Errorf("查询麦位状态失败: %v", err)
	}

	// 4. 构建麦位信息，关联查询用户信息
	var micInfos []*__.MicInfo
	for _, mic := range roomMics {
		micInfo := &__.MicInfo{
			Position: int32(mic.MicPosition),
			Status:   int32(mic.Status),
			IsLocked: mic.IsLocked == 1,
		}

		// 如果麦位被占用，查询用户信息
		if mic.Status == 1 && mic.UserId > 0 {
			var user models.User
			err = global.DB.Where("id = ?", mic.UserId).First(&user).Error
			if err == nil {
				micInfo.UserId = mic.UserId
				micInfo.UserNickname = user.Nickname
				micInfo.UserAvatar = user.Avatar
				if !mic.OccupyTime.IsZero() {
					micInfo.OccupyTime = mic.OccupyTime.Format("2006-01-02 15:04:05")
				}

				// 检查用户是否被禁言
				var roomMember models.RoomMember
				err = global.DB.Where("user_id = ? AND room_id = ?", mic.UserId, req.RoomId).First(&roomMember).Error
				if err == nil && roomMember.IsMuted == 1 {
					// 检查禁言是否已过期
					if roomMember.MuteEndTime.IsZero() || time.Now().Before(roomMember.MuteEndTime) {
						micInfo.IsMuted = true
					} else {
						// 禁言已过期，自动解除
						global.DB.Model(&roomMember).Updates(map[string]interface{}{
							"is_muted":      0,
							"mute_end_time": nil,
						})
						micInfo.IsMuted = false
					}
				} else {
					micInfo.IsMuted = false
				}
			}
		}

		micInfos = append(micInfos, micInfo)
	}

	// 5. 获取待审批的申请列表
	var applications []models.MicApplication
	err = global.DB.Where("room_id = ? AND status = 0", req.RoomId).Order("apply_time").Find(&applications).Error
	if err != nil {
		return nil, fmt.Errorf("查询申请列表失败: %v", err)
	}

	// 6. 构建申请信息，关联查询用户信息
	var pendingApplications []*__.MicApplication
	for _, app := range applications {
		// 验证申请用户是否仍在房间中
		var userRoom models.UserRoom
		err = global.DB.Where("user_id = ? AND room_id = ? AND is_online = 1", app.UserId, req.RoomId).First(&userRoom).Error
		if err != nil {
			// 用户已离开房间，自动取消申请
			global.DB.Model(&app).Updates(map[string]interface{}{
				"status":      3, // 已取消
				"handle_time": time.Now(),
				"reason":      "用户已离开房间",
			})
			continue
		}

		// 查询申请用户信息
		var user models.User
		err = global.DB.Where("id = ?", app.UserId).First(&user).Error
		if err != nil {
			continue // 跳过无效用户
		}

		appInfo := &__.MicApplication{
			Id:           app.Id,
			UserId:       app.UserId,
			UserNickname: user.Nickname,
			UserAvatar:   user.Avatar,
			ApplyTime:    app.ApplyTime.Format("2006-01-02 15:04:05"),
			Status:       int32(app.Status),
		}

		pendingApplications = append(pendingApplications, appInfo)
	}

	// 7. 返回格式化的麦位和申请信息
	return &__.GetMicStatusResp{
		Mics:                micInfos,
		PendingApplications: pendingApplications,
	}, nil
}
