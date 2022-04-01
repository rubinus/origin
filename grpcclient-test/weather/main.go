package main

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/origin/configs"
  "github.com/gitcpu-io/origin/grpcclients"
  pb_weather "github.com/gitcpu-io/origin/pb/weather"
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
  grpcclients.RPCClientsRun(nil)
  time.Sleep(1 * time.Second)

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  hReq := &pb_weather.ListRequest{City: "深圳市"}
  if response, err := grpcclients.RpcWeather(ctx, hReq); response != nil {
    bytes, _ := zgo.Utils.Marshal(response)
    fmt.Printf("RpcWeather: %s \n\n", string(bytes))
  } else {
    zgo.Log.Error(err)
  }

}
