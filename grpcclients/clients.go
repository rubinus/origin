package grpcclients

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/origin/configs"
  "github.com/gitcpu-io/origin/pb/helloworld"
  pb_weather "github.com/gitcpu-io/origin/pb/weather"
  "github.com/gitcpu-io/zgo"
  "time"
)

/*
@Time : 2019-08-31 15:17
@Author : rubinus.chu
@File : clients
@project: origin
*/

// HelloWorldClient 可以起名为你的 xxxxClient
var HelloWorldClient pb_helloworld.HelloWorldServiceClient

// YourClient here
//var YourClient pb_your.YourServiceClient

// WeatherClient 可以起名为你的WeatherClient
var WeatherClient pb_weather.WeatherServiceClient

func RPCClientsRun(ch chan string) {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  if configs.Conf.StartService &&
    configs.Conf.StartServiceDiscover { //使用服务发现模式
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
    //正常模式
    go helloWorldClient(ctx, configs.Conf.RpcHost, configs.Conf.RpcPort)

    go weatherClient(ctx, configs.Conf.RpcHost, configs.Conf.RpcPort)

    //go yourClient(ctx, "your call rpc host", "your call rpc port")
    //请在下面逐个添加你的proto生成的pb的client

  }
}

// helloWorldClient 客户端封装
func helloWorldClient(ctx context.Context, address string, port int) {
  conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure(),zgo.Grpc.WithCallOptions(1024 * 1024 * 16))
  if err != nil {
    errStr := fmt.Sprintf("helloWorldClient timeout, Host: %s, Port: %d", address, port)
    zgo.Log.Error(errStr)
    return
  }
  client := pb_helloworld.NewHelloWorldServiceClient(conn)
  HelloWorldClient = client
}

// yourClient 自己改名
//func yourClient(ctx context.Context, address, port string) {
//	conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure(),zgo.Grpc.WithCallOptions(1024 * 1024 * 16))
//	if err != nil {
//		errStr := fmt.Sprintf("yourClient timeout, Host: %s, Port: %s", address, port)
//		zgo.Log.Error(errStr)
//		return
//	}
//	client := pb_your.NewYourServiceClient(conn)
//	YourClient = client
//}

// weatherClient 客户端封装
func weatherClient(ctx context.Context, address string, port int) {
  conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure(),zgo.Grpc.WithCallOptions(1024 * 1024 * 16))
  if err != nil {
    errStr := fmt.Sprintf("weatherClient timeout, Host: %s, Port: %d", address, port)
    zgo.Log.Error(errStr)
    return
  }
  client := pb_weather.NewWeatherServiceClient(conn)
  WeatherClient = client
}
