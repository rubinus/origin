/*
@Time : 2019-03-02 11:40
@Author : zhangjianguo
@File : main
@Software: GoLand
*/

// Package main implements a grpcserver for Greeter service.
package main

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/origin/configs"
	"github.com/gitcpu-io/origin/examples/demo_grpc/go/handler"
	pb "github.com/gitcpu-io/origin/examples/demo_grpc/go/pb"
	"github.com/gitcpu-io/zgo"
)

func main() {
	configs.InitConfig("", "local", "", "", 0, 0)

	err := zgo.Engine(&zgo.Options{
		Env:      configs.Conf.Env,
		Loglevel: configs.Conf.Loglevel,
		Project:  configs.Conf.Project,
	})
	if err != nil {
		panic(err)
	}

	server, _ := zgo.Grpc.Server(context.TODO())

	pb.RegisterGreeterServer(server, &handler.Server{})

	msg, _ := zgo.Grpc.Run(context.TODO(), server, 50051)

	fmt.Println(msg)
}
