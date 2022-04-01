package main

import (
	"context"
	"fmt"
	"github.com/gitcpu-io/origin/configs"
	"github.com/gitcpu-io/origin/grpcclient"
	"github.com/gitcpu-io/origin/pb/helloworld"
	"github.com/gitcpu-io/zgo"
	"time"
)

func main() {
	configs.InitConfig("", "local", "", "", 0, 0)

	err := zgo.Engine(&zgo.Options{
		CPath:    configs.Conf.CPath,
		Env:      configs.Conf.Env,
		Loglevel: configs.Conf.Loglevel,
		Project:  configs.Conf.Project,
	})
	if err != nil {
		zgo.Log.Error(err)
		return
	}

	//start grpc clients
	grpcclient.RPCClientsRun(nil)
	time.Sleep(1 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//组织请求参数
	request1 := &pb_helloworld.HelloRequest_Request{
		Url:   "example.com",
		Title: "origin title 101",
		Ins:   []string{"example.com", "origin 101"},
	}
	request2 := &pb_helloworld.HelloRequest_Request{
		Url:   "example.com",
		Title: "origin title 102",
		Ins:   []string{"example.com", "origin 102"},
	}
	var requests []*pb_helloworld.HelloRequest_Request
	requests = append(requests, request1)
	requests = append(requests, request2)

	hReq := &pb_helloworld.HelloRequest{Name: "origin project hello", Age: 30, Requests: requests}
	if response, err := grpcclient.RpcHelloWorld(ctx, hReq); response != nil {
		bytes, _ := zgo.Utils.Marshal(response)
		fmt.Printf("RpcHelloWorld: %s \n\n", string(bytes))
	} else {
		zgo.Log.Error(err)
	}

}
