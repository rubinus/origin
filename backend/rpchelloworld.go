package backend

import (
	"context"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"git.zhugefang.com/goymd/visource/pb"
)

/*
@Time : 2019-06-15 11:09
@Author : rubinus.chu
@File : rpchelloworld
@project: visource
*/

func RpcHelloWorld(ctx context.Context, address, port string, request *pb.HelloRequest) *pb.HelloReply {
	out := make(chan *pb.HelloReply)
	go func() {
		conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure())
		if err != nil {
			zgo.Log.Error(err)
			return
		}
		defer conn.Close()
		client := pb.NewGreeterClient(conn)
		response, err := client.SayHello(ctx, request)
		if err != nil {
			zgo.Log.Error(err)
			return
		}
		out <- response
	}()

	select {
	case <-ctx.Done():
		errStr := fmt.Sprintf("RpcHelloWorld timeout, Host: %s, Port: %s", address, port)
		zgo.Log.Error(errStr)
		return nil
	case r := <-out:
		return r
	}

}
