package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/user/basic/config"
	inits2 "server/user/basic/inits"
	"server/user/handler"
	__ "server/user/proto"
)

func main() {
	config.Inits()
	inits2.InitMysql()
	inits2.RedisInit()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8889))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterUserServer(s, &handler.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
