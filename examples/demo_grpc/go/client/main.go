/*
@Time : 2019-03-02 11:42
@Author : zhangjianguo
@File : main
@Software: GoLand
*/

package main

import (
	"context"
	"github.com/gitcpu-io/origin/configs"
	"github.com/gitcpu-io/zgo"
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
	configs.InitConfig("", "local", "", "", 0, 0)

	err := zgo.Engine(&zgo.Options{
		Env:      configs.Conf.Env,
		Loglevel: configs.Conf.Loglevel,
		Project:  configs.Conf.Project,
	})
	if err != nil {
		panic(err)
	}
	conn, err := zgo.Grpc.Client(context.TODO(), address, 50051, zgo.Grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the grpcserver and print out its response.
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
