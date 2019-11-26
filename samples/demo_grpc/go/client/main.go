/*
@Time : 2019-03-02 11:42
@Author : zhangjianguo
@File : main
@Software: GoLand
*/

package main

import (
	"context"
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gocore/zgo"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost"
	defaultName = "world"
)

func main() {
	config.InitConfig("local", "", "", "", "")

	err := zgo.Engine(&zgo.Options{
		Env:      config.Conf.Env,
		Loglevel: config.Conf.Loglevel,
		Project:  config.Conf.Project,
	})
	conn, err := zgo.Grpc.Client(context.TODO(), address, "", zgo.Grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
