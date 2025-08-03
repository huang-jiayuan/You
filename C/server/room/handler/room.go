package handler

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"server/models"
	"server/pkg"
	"server/room/basic/global"
	__ "server/room/proto"
)

type Server struct {
	__.UnimplementedRoomServer
}

// 示例
func (s *Server) Stream(_ context.Context, in *__.StreamReq) (*__.StreamResp, error) {
	return &__.StreamResp{}, nil
}

func (s *Server) CreateRoom(_ context.Context, in *__.CreateRoomStreamReq) (*__.CreateRoomStreamResp, error) {
	// 查询用户存不存在
	var user models.User
	err := global.DB.Where("id = ?", in.UserId).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	// 检查用户状态
	if user.Status != 0 {
		return nil, fmt.Errorf("账号状态异常，无法创建房间")
	}

	// 如果在没有实名的状态下
	if user.IsAuth == 0 {
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
	// 统计用户活跃房间数量
	err = global.DB.Model(&models.Room{}).Where("user_id = ? AND (closed_at IS NULL) AND status != '1'", in.UserId).Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("查询用户房间信息失败: %v", err)
	}
	// 如果大于0 说明有活跃房间
	if count > 0 {
		var activeRoom models.Room
		//查询如果关闭时间为空并且房间状态为1已关闭 说明没有开着的房间了
		err = global.DB.Where("user_id = ? AND closed_at IS NULL AND status != '1'", in.UserId).First(&activeRoom).Error
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
	var user models.User
	err := global.DB.Where("status = ?", in.Status).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("查询失败")
	}
	// 检查用户状态
	if user.Status != 0 {
		return nil, fmt.Errorf("账号状态异常，无法修改房间信息")
	}
	var room models.Room
	err = global.DB.Where("id = ?", in.Id).Find(&room).Error
	if err != nil {
		return nil, fmt.Errorf("房间查询失败: %v", err)
	}
	//检查房间在不在
	if room.Id == 0 {
		return nil, fmt.Errorf("房间不存在或异常")
	}
	update := models.Room{
		Id:           int64(in.Id),
		Announcement: in.Announcement,
	}
	//只有不为空才检测敏感词
	if in.Announcement != "" {
		//调用敏感词检测
		isClean := pkg.Textcen(in.Announcement)
		if !isClean {
			return nil, fmt.Errorf("房间公告包含敏感信息，请修改后重试")
		}
	} else {
		fmt.Printf("房间公告为空，跳过敏感词检测\n")
	}
	//修改
	err = global.DB.Updates(update).Error
	if err != nil {
		return nil, fmt.Errorf("修改失败: %v", err)
	}
	return &__.UpdateRoomStreamResp{
		Id: int32(update.Id),
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
	//房间状态正常并且关闭时间为空
	x := "status = '0' AND closed_at IS NULL"
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
	//根据标签id查 并且房间状态正常
	c := "status = '0' AND closed_at IS NULL AND tag_id = ?"
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
	//状态为正常并且房间关闭时间为空的情况下
	z := "status = '0' AND closed_at IS NULL AND room_name LIKE ?"
	// 模糊查询
	searchPattern := "%" + req.Keyword + "%"
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
	err = global.DB.Scopes(paginateScope).Where(z, searchPattern).Order("fk_member_room DESC").
		Find(&rooms).Error
	if err != nil {
		return nil, fmt.Errorf("搜索房间失败: %v", err)
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
	return &__.SearchRoomsResp{
		Rooms:    roomInfos,
		Total:    int32(total),
		Page:     page,
		PageSize: pageSize,
		Keyword:  req.Keyword,
	}, nil
}

// 关闭房间
func (s *Server) CloseRoom(_ context.Context, in *__.CloseRoomStreamReq) (*__.CloseRoomStreamResp, error) {
	return &__.CloseRoomStreamResp{}, nil
}
