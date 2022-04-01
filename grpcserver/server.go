package grpcserver

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/origin/configs"
  "github.com/gitcpu-io/origin/grpchandlers"
  "github.com/gitcpu-io/origin/pb/helloworld"
  pb_weather "github.com/gitcpu-io/origin/pb/weather"
  "github.com/gitcpu-io/zgo"
  "math"
  "time"
)

/*
@Time : 2019-06-10 18:47
@Author : rubinus.chu
@File : server
@project: origin
*/

func Start() {
  ka := time.Duration(math.MaxInt64)
  server, err := zgo.Grpc.Server(context.TODO(),zgo.Grpc.WithKeepalive(ka,ka,ka, 5 * time.Second,20 * time.Second))
  if err != nil {
    fmt.Println("Grpc server is error :", err)
    return
  }

  pb_helloworld.RegisterHelloWorldServiceServer(server, &grpchandlers.HelloWorldServer{})

  //在这里继续添加你的rpc服务
  // todo add your rpc service
  pb_weather.RegisterWeatherServiceServer(server, &grpchandlers.WeatherServer{})


  fmt.Printf("Now listening GRPC Server on: %d\n", configs.Conf.RpcPort)
  _, err = zgo.Grpc.Run(context.TODO(), server, configs.Conf.RpcPort)
  if err != nil {
    fmt.Println("Grpc server is error :", err)
    return
  }
}
