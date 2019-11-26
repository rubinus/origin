/*
@Time : 2019-03-02 11:40
@Author : zhangjianguo
@File : main
@Software: GoLand
*/

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gobase/origin/samples/demo_grpc/go/handler"
	pb "git.zhugefang.com/gobase/origin/samples/demo_grpc/go/pb"
	"git.zhugefang.com/gocore/zgo"
)

func main() {
	config.InitConfig("local", "", "", "", "")

	zgo.Engine(&zgo.Options{
		Env:      config.Conf.Env,
		Loglevel: config.Conf.Loglevel,
		Project:  config.Conf.Project,
	})

	server, _ := zgo.Grpc.Server(context.TODO())

	pb.RegisterGreeterServer(server, &handler.Server{})

	msg, _ := zgo.Grpc.Run(context.TODO(), server, "")

	fmt.Println(msg)
}
