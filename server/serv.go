package server

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"git.zhugefang.com/goymd/visource/config"
	"git.zhugefang.com/goymd/visource/grpchandlers"
	"git.zhugefang.com/goymd/visource/pb/helloworld"
)

/*
@Time : 2019-06-10 18:47
@Author : rubinus.chu
@File : serv
@project: account
*/

func Start() {
	server, err := zgo.Grpc.Server(context.TODO())
	if err != nil {
		fmt.Println("-----grpc server is error :", err)
		return
	}

	pb_helloworld.RegisterHelloWorldServiceServer(server, &grpchandlers.HelloWorldServer{})

	msg, err := zgo.Grpc.Run(context.TODO(), server, config.Conf.RpcPort)
	if err != nil {
		fmt.Println("=====grpc server is error :", err)
		return
	}

	fmt.Println(msg)
}
