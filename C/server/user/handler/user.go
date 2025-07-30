package handler

import (
	"context"
	__ "server/user/proto"
)

type Server struct {
	__.UnimplementedUserServer
}

//示例
func (s *Server) Stream(_ context.Context, in *__.StreamReq) (*__.StreamResp, error) {
	return &__.StreamResp{}, nil
}
