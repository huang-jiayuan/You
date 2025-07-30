package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/room/basic/config"
	inits2 "server/room/basic/inits"
	"server/room/handler"
	__ "server/room/proto"
)

func main() {
	config.Inits()
	inits2.InitMysql()
	inits2.RedisInit()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterRoomServer(s, &handler.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
