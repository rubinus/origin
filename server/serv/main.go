package main

/*
@Time : 2019-06-10 18:01
@Author : rubinus.chu
@File : main
@project: account
*/

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"git.zhugefang.com/goymd/visource/config"
	"git.zhugefang.com/goymd/visource/grpchandlers"
	"git.zhugefang.com/goymd/visource/pb"
)

func main() {
	config.InitConfig("local", "", "", "")
	zgo.Engine(&zgo.Options{
		Env:      config.Conf.Env,
		Loglevel: config.Conf.Loglevel,
		Project:  config.Conf.Project,
	})

	server, _ := zgo.Grpc.Server(context.TODO())

	pb.RegisterGreeterServer(server, &grpchandlers.Server{})

	msg, _ := zgo.Grpc.Run(context.TODO(), server)

	fmt.Println(msg)
}
