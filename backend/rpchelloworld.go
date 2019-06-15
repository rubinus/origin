package backend

import (
	"context"
	"git.zhugefang.com/gocore/zgo"
	"git.zhugefang.com/goymd/visource/pb"
)

/*
@Time : 2019-06-15 11:09
@Author : rubinus.chu
@File : rpchelloworld
@project: visource
*/

func RpcHelloWorld(ctx context.Context, address, port string, request *pb.HelloRequest) chan *pb.HelloReply {
	out := make(chan *pb.HelloReply)
	go func() {

		conn, err := zgo.Grpc.Client(ctx, address, port, zgo.Grpc.WithInsecure())
		if err != nil {
			zgo.Log.Errorf("did not connect: %v", err)
			out <- nil
			return
		}
		defer conn.Close()

		client := pb.NewGreeterClient(conn)

		//--------SignRequest-------发起签到rpc调用
		response, err := client.SayHello(ctx, request)

		if err != nil {
			zgo.Log.Error(err)
			out <- nil
			return
		}

		select {
		case <-ctx.Done():
			zgo.Log.Errorf("timeout call GRPC server %s, port %s", address, port)
			out <- nil
			return
		default:
			out <- response
		}

	}()
	return out
}
