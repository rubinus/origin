package main

import (
  "context"
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/origin/grpcclients"
  "github.com/gitcpu-io/origin/pb/helloworld"
  "github.com/gitcpu-io/zgo"
  "os"
  "time"
)

func main() {
  config.InitConfig("local", "", "", "", "")

  getwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  err = zgo.Engine(&zgo.Options{
    CPath: fmt.Sprintf("%s/%s",getwd,"config"),
    Env:      config.Conf.Env,
    Loglevel: config.Conf.Loglevel,
    Project:  config.Conf.Project,
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
  if response, err := grpcclients.RpcHelloWorld(ctx, hReq); response != nil {
    bytes, _ := zgo.Utils.Marshal(response)
    fmt.Printf("RpcHelloWorld: %s \n\n", string(bytes))
  } else {
    zgo.Log.Error(err)
  }

}
