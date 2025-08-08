package handler

import (
	"context"
	"fmt"
	"server/models"
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

func (s *Server) SendGifts(_ context.Context, in *__.SendGiftsReq) (*__.SendGiftsResp, error) {
	user := &models.User{}
	_, err := user.FindUserById(in.SendUserId)
	if err != nil {
		fmt.Println("未查询到当前送礼的用户")
		return nil, err
	}
	_, err = user.FindUserById(in.ReceiveUserId)
	if err != nil {
		fmt.Println("未查询到当前接受礼物的用户")
		return nil, err
	}
	room := &models.Room{}
	err = room.GetFindRoomById(in.RoomId)
	if err != nil {
		fmt.Println("未查询到该房间号")
		return nil, err
	}
	info := &models.GiftInfo{}
	err = info.GetFindGiftById(in.GiftId)
	if err != nil {
		fmt.Println("未查询到当前的礼物")
		return nil, err
	}
	tx := global.DB.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()
	giftprice := info.Price * float64(in.SendCount)
	if user.Balance < giftprice {
		fmt.Println("用户余额不够购买此礼物")
		tx.Rollback()
		return nil, err
	}
	record := &models.GiftSendRecord{
		SendUserId:    uint64(in.SendUserId),
		ReceiveUserId: uint64(in.ReceiveUserId),
		RoomId:        in.RoomId,
		GiftId:        in.GiftId,
		SendCount:     int32(in.SendCount),
		TotalDiamond:  int32(giftprice),
		TotalPrice:    giftprice,
		SendType:      in.SendType,
		Message:       in.Message,
		Status:        in.Status,
		ClientIp:      in.ClientIp,
	}
	err = global.DB.Create(&record).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("刷礼物记录失败")
		return nil, err
	}
	tx.Commit()
	return &__.SendGiftsResp{Greet: "刷礼物成功"}, nil
}
