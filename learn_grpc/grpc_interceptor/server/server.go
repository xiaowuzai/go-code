package main

import (
	"context"
	"fmt"
	"net"

	"github.com/xiaowuzai/go-code/learn_grpc/grpc_metadata/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func timeInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("接收到一个请求")
	res, err := handler(ctx, req)
	fmt.Println("请求结束")
	return res, err
}

func main() {

	gs := grpc.NewServer(grpc.UnaryInterceptor(timeInterceptor))
	defer gs.GracefulStop()

	proto.RegisterGreeterServer(gs, &Server{})

	lis, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("error listening")
	}

	err = gs.Serve(lis)
	if err != nil {
		panic("error serve")
	}
}

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {

	md, has := metadata.FromIncomingContext(ctx)
	if !has {
		fmt.Println("get metadata from context failed")
		return nil, nil
	}

	for k, v := range md {
		fmt.Println(k, v)
	}
	return &proto.HelloReply{
		Message: "Hello " + req.Name,
	}, nil
}
