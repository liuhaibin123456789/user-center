package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user-center/global"
	"user-center/proto"
	"user-center/service"
	"user-center/tool"
)

func main() {

	err := tool.LinkMysql()
	if err != nil {
		log.Println(err)
		return
	}

	err = tool.InitRedis()
	if err != nil {
		log.Println(err)
		return
	}

	//todo 提供文件服务，传输proto文件和端口，应该使用websocket进行双向传输

	listener, err := net.Listen("tcp", global.Port)
	if err != nil {
		log.Println(err)
		return
	}
	//开启grpc服务器
	server := grpc.NewServer()
	//将服务注册绑定到grpc服务器
	proto.RegisterUserCenterServer(server, service.Service{})
	log.Println("grpc listen...")
	err = server.Serve(listener)
	if err != nil {
		log.Println(err)
		return
	}
}
