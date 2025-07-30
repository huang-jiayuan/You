package handler

import (
	"context"
	__ "server/room/proto"
)

type Server struct {
	__.UnimplementedRoomServer
}

// 示例
func (s *Server) Stream(_ context.Context, in *__.StreamReq) (*__.StreamResp, error) {
	return &__.StreamResp{}, nil
}
