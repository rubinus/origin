package backend

import (
	"context"
	"fmt"
	"git.zhugefang.com/gobase/origin/config"
	"git.zhugefang.com/gobase/origin/pb/helloworld"
	"git.zhugefang.com/gocore/zgo"
	"time"
)

/*
@Time : 2019-08-31 15:17
@Author : rubinus.chu
@File : rpc_clients
@project: origin
*/

// 可以起名为你的 xxxxClient
var HelloworldClient pb_helloworld.HelloWorldServiceClient

// you are client
//var YourClient pb_your.YourServiceClient

func RPCClientsRun() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//create client	//请在下面逐个添加你的proto生成的pb的client
	go helloWorldClient(ctx, config.Conf.RpcHost, config.Conf.RpcPort)

	//go yourClient(ctx, "your call rpc host", "your call rpc port")
}

func helloWorldClient(ctx context.Context, address, port string) {
	conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure())
	if err != nil {
		errStr := fmt.Sprintf("helloWorldClient timeout, Host: %s, Port: %s", address, port)
		zgo.Log.Error(errStr)
		return
	}
	client := pb_helloworld.NewHelloWorldServiceClient(conn)
	HelloworldClient = client
}

// yourClient 自己改名
//func yourClient(ctx context.Context, address, port string) {
//	conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure())
//	if err != nil {
//		errStr := fmt.Sprintf("yourClient timeout, Host: %s, Port: %s", address, port)
//		zgo.Log.Error(errStr)
//		return
//	}
//	client := pb_your.NewYourServiceClient(conn)
//	YourClient = client
//}
