/*
@Time : 2019-03-02 11:51
@Author : rubinus
@File : hellworld
@Software: origin
*/
package grpchandlers

import (
	"context"
	"fmt"
	"github.com/rubinus/origin/pb/helloworld"
	"log"
)

// 可以起名为你的 xxxxServer
type HelloWorldServer struct{}

// SayHello implements func
func (s *HelloWorldServer) SayHello(ctx context.Context, request *pb_helloworld.HelloRequest) (*pb_helloworld.HelloResponse, error) {
	log.Printf("Received: Name %v", request.Name)
	log.Printf("Received: Age %v", request.Age)
	for _, v := range request.Requests {
		fmt.Printf("Request: Url:%s, Title:%s, 字符串[]:%s \n", v.Url, v.Title, v.Ins)
	}

	//构建[]
	var infos []*pb_helloworld.Info
	for i := 0; i < 5; i++ {
		obj := &pb_helloworld.Info{
			Email: fmt.Sprintf("%s%d%s", "test", i, "@zhuge.com"),
			Money: float64(1000 * i),
		}
		infos = append(infos, obj)
	}

	//构建Map
	var tmpMap = make(map[string]int32)
	tmpMap["K1"] = 100
	tmpMap["K2"] = 200
	tmpMap["K3"] = 300

	//构建自定义的kv
	var myMap = make(map[string]*pb_helloworld.Info)
	myMap["my1"] = &pb_helloworld.Info{Email: "my1@zhuge.com", Money: 1000.00}
	myMap["my2"] = &pb_helloworld.Info{Email: "my2@zhuge.com", Money: 2000.00}

	return &pb_helloworld.HelloResponse{
		Message:  "Hello " + request.Name, //string
		Infos:    infos,                   //数组
		Ans:      []string{"a", "b", "this is a string"},
		Projects: tmpMap,
		MyMap:    myMap,
	}, nil
}
