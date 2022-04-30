package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user-center/global"
	"user-center/proto"
	"user-center/service"
)

func main() {
	listener, err := net.Listen("tcp", global.Port)
	if err != nil {
		log.Println(err)
		return
	}
	//开启grpc服务器
	server := grpc.NewServer()
	//将服务注册绑定到grpc服务器
	proto.RegisterUserCenterServer(server, service.Service{})
	err = server.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}
}
