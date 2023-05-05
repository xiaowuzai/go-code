package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/xiaowuzai/go-code/learn_grpc/grpc_http/proto"
)

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
