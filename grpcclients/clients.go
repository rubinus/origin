package grpcclients

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/origin/pb/helloworld"
  "github.com/gitcpu-io/zgo"
  "time"
)

/*
@Time : 2019-08-31 15:17
@Author : rubinus.chu
@File : clients
@project: origin
*/

// HelloworldClient 可以起名为你的 xxxxClient
var HelloworldClient pb_helloworld.HelloWorldServiceClient

// you are client
//var YourClient pb_your.YourServiceClient

func RPCClientsRun(ch chan string) {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  if config.Conf.StartService == true &&
    config.Conf.StartServiceDiscover == true { //使用服务发现模式
    go func() {
      for value := range ch {
        lbRes, err := zgo.Service.LB(value) //变化的服务
        if err != nil {
          zgo.Log.Error(fmt.Sprintf("%s 服务取Grpc负载,", value), err)
          continue
        }

        switch value {
        case "origin.bffp": //自己做为客户端连接自己的服务端测试
          //create client	//请在下面逐个添加你的proto生成的pb的client
          go helloWorldClient(ctx, lbRes.SvcHost, lbRes.SvcGrpcPort)

        case "other":
          //继续通过服务名，来再次初始化host grpc port
        }

        zgo.Log.Warnf("监听到Grpc服务：%s,正在使用负载节点 Host: %s, grpc_host: %s", value, lbRes.SvcHost, lbRes.SvcGrpcPort)

      }
    }()

  } else {

    go helloWorldClient(ctx, config.Conf.RpcHost, config.Conf.RpcPort)

  }

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
