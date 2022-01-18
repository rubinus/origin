package grpcclients

import (
  "context"
  "errors"
  "fmt"
  "github.com/gitcpu-io/origin/pb/helloworld"
)

/*
@Time : 2019-06-15 11:09
@Author : rubinus.chu
@File : helloworld
@project: origin
*/

func RpcHelloWorld(ctx context.Context, request *pb_helloworld.HelloRequest) (*pb_helloworld.HelloResponse, error) {
  out := make(chan *pb_helloworld.HelloResponse)
  errCh := make(chan error)
  go func() {
    if HelloworldClient == nil {
      errCh <- errors.New("HelloworldClient not ready")
      return
    }
    response, err := HelloworldClient.SayHello(ctx, request)
    if err != nil {
      errCh <- err
      return
    }
    out <- response
  }()

  select {
  case <-ctx.Done():
    errStr := fmt.Sprintf("RpcHelloWorld timeout")
    return nil, errors.New(errStr)

  case err := <-errCh:
    return nil, err

  case r := <-out:
    return r, nil
  }

}
