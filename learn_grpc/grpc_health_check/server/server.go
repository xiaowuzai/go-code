package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/xiaowuzai/go-code/learn_grpc/grpc_metadata/proto"
)

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("接收到一个请求")

	md, has := metadata.FromIncomingContext(ctx)
	if !has {
		return nil, status.Errorf(codes.Unauthenticated, "Authentication is empty")
	}

	appid, has := md["appid"]
	if !has {
		return nil, status.Errorf(codes.Unauthenticated, "Authentication error")
	}
	appkey, has := md["appkey"]
	if !has {
		return nil, status.Errorf(codes.Unauthenticated, "Authentication error")
	}

	log.Printf("appId = %s, appKey = %s\n", appid[0], appkey[0])

	res, err := handler(ctx, req)
	fmt.Println("请求结束")
	return res, err
}

func main() {

	// New grpc server
	// gs := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	gs := grpc.NewServer()
	defer gs.GracefulStop()

	// Register Server
	proto.RegisterGreeterServer(gs, &Server{})

	// Register Health server
	grpc_health_v1.RegisterHealthServer(gs, health.NewServer())

	log.Printf("grpc server listening on %s", "8088")
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
