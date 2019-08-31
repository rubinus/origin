package backend

import (
	"context"
	"errors"
	"fmt"
	"git.zhugefang.com/gobase/origin/pb/helloworld"
)

/*
@Time : 2019-06-15 11:09
@Author : rubinus.chu
@File : rpc_helloworld
@project: origin
*/

func RpcHelloWorld(ctx context.Context, request *pb_helloworld.HelloRequest) (*pb_helloworld.HelloResponse, error) {
	out := make(chan *pb_helloworld.HelloResponse)
	errCh := make(chan error)
	go func() {
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
