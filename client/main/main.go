package main

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"git.zhugefang.com/goymd/visource/config"
	"git.zhugefang.com/goymd/visource/pb"
	"log"
	"time"
)

const (
	address     = "localhost"
	defaultName = "world"
)

func main() {
	config.InitConfig("dev", "", "", "")

	err := zgo.Engine(&zgo.Options{
		Env:      config.Conf.Env,
		Loglevel: config.Conf.Loglevel,
		Project:  config.Conf.Project,
	})
	conn, err := zgo.Grpc.Client(context.TODO(), address, config.Conf.RpcPort, zgo.Grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//组织请求参数
	request1 := &pb.HelloRequest_Request{
		Url:   "zhuge.com",
		Title: "诸葛找房111111",
		Ins:   []string{"zhuge.com", "诸葛找房111111"},
	}
	request2 := &pb.HelloRequest_Request{
		Url:   "zhuge.com",
		Title: "诸葛找房222222",
		Ins:   []string{"zhuge.com", "诸葛找房222222"},
	}
	var requests []*pb.HelloRequest_Request
	requests = append(requests, request1)
	requests = append(requests, request2)

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName, Age: 30, Requests: requests})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	b, _ := zgo.Utils.Marshal(r)

	fmt.Println(string(b))

}
