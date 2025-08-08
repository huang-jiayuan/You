package handler

import (
	"context"
	"server/models"
	__ "server/room/proto"
	"server/user/basic/global"
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
