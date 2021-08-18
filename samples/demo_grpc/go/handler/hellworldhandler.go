/*
@Time : 2019-03-02 11:51
@Author : zhangjianguo
@File : hellworldHandler
@Software: GoLand
*/
package handler

import (
	"context"
	pb "github.com/rubinus/origin/samples/demo_grpc/go/pb"
	"log"
)

// grpcserver is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
