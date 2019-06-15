package main

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"git.zhugefang.com/goymd/visource/backend"
	"git.zhugefang.com/goymd/visource/config"
	"git.zhugefang.com/goymd/visource/pb"
	"time"
)

func main() {
	config.InitConfig("dev", "", "", "")

	err := zgo.Engine(&zgo.Options{
		Env:      config.Conf.Env,
		Loglevel: config.Conf.Loglevel,
		Project:  config.Conf.Project,
	})
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	hReq := &pb.HelloRequest{Name: "hello", Age: 30, Requests: requests}
	if response, _ := backend.RpcHelloWorld(ctx, config.Conf.RpcHost, config.Conf.RpcPort, hReq); response != nil {
		bytes, _ := zgo.Utils.Marshal(response)
		fmt.Printf("RpcHelloWorld: %s \n\n", string(bytes))
	}

}
