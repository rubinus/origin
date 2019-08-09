package backend

import (
	"context"
	"errors"
	"fmt"
	"git.zhugefang.com/gobase/base-to-base-wait-copy/pb/helloworld"
	"git.zhugefang.com/gocore/zgo"
)

/*
@Time : 2019-06-15 11:09
@Author : rubinus.chu
@File : rpchelloworld
@project: base-to-base-wait-copy
*/

func RpcHelloWorld(ctx context.Context, address, port string, request *pb_helloworld.HelloRequest) (*pb_helloworld.HelloResponse, error) {
	out := make(chan *pb_helloworld.HelloResponse)
	errCh := make(chan error)
	go func() {
		conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure())
		if err != nil {
			//zgo.Log.Error(err)
			errCh <- err
			return
		}
		defer conn.Close()
		client := pb_helloworld.NewHelloWorldServiceClient(conn)
		response, err := client.SayHello(ctx, request)
		if err != nil {
			//zgo.Log.Error(err)
			errCh <- err
			return
		}
		out <- response
	}()

	select {
	case <-ctx.Done():
		errStr := fmt.Sprintf("RpcHelloWorld timeout, Host: %s, Port: %s", address, port)
		//zgo.Log.Error(errStr)
		return nil, errors.New(errStr)

	case err := <-errCh:
		return nil, err

	case r := <-out:
		return r, nil
	}

}
