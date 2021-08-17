package backend

import (
	"context"
	"errors"
	"fmt"
	"github.com/rubinus/origin/pb/helloworld"
	"github.com/rubinus/zgo"
	"time"
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

func CallRpcHelloworld() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//组织请求参数
	request1 := &pb_helloworld.HelloRequest_Request{
		Url:   "zhuge.com",
		Title: "诸葛找房111111",
		Ins:   []string{"zhuge.com", "诸葛找房111111"},
	}
	request2 := &pb_helloworld.HelloRequest_Request{
		Url:   "zhuge.com",
		Title: "诸葛找房222222",
		Ins:   []string{"zhuge.com", "诸葛找房222222"},
	}
	var requests []*pb_helloworld.HelloRequest_Request
	requests = append(requests, request1)
	requests = append(requests, request2)

	hReq := &pb_helloworld.HelloRequest{Name: "origin project hello", Age: 30, Requests: requests}
	if response, err := RpcHelloWorld(ctx, hReq); response != nil {
		bytes, _ := zgo.Utils.Marshal(response)
		fmt.Printf("RpcHelloWorld: %s \n\n", string(bytes))
	} else {
		zgo.Log.Error(err)
	}
}
