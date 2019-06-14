/*
@Time : 2019-03-02 11:51
@Author : rubinus
@File : hellworldHandler
@Software: visource
*/
package grpchandlers

import (
	"context"
	"fmt"
	"git.zhugefang.com/goymd/visource/pb"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: Name %v", in.Name)
	log.Printf("Received: Age %v", in.Age)
	for _, v := range in.Requests {
		fmt.Printf("Request: Url:%s, Title:%s, 字符串[]:%s \n", v.Url, v.Title, v.Ins)
	}

	//构建[]
	var infos []*pb.Info
	for i := 0; i < 5; i++ {
		obj := &pb.Info{
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
	var myMap = make(map[string]*pb.Info)
	myMap["my1"] = &pb.Info{Email: "my1@zhuge.com", Money: 1000.00}
	myMap["my2"] = &pb.Info{Email: "my2@zhuge.com", Money: 2000.00}

	return &pb.HelloReply{
		Message:  "Hello " + in.Name, //string
		Infos:    infos,              //数组
		Ans:      []string{"a", "b", "this is a string"},
		Projects: tmpMap,
		MyMap:    myMap,
		Code:     200,
	}, nil
}
