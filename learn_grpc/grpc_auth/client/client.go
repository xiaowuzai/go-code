package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/xiaowuzai/go-code/learn_grpc/grpc_metadata/proto"
)

func authInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	md := metadata.New(map[string]string{
		"appid":  "123123",
		"appkey": "123123appkey",
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)

}

type credential struct{}

func (c *credential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "123123",
		"appkey": "123123appkey",
	}, nil
}
func (c *credential) RequireTransportSecurity() bool {
	return false
}

func main() {

	// opt := grpc.WithUnaryInterceptor(authInterceptor)
	grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("127.0.0.1:8088",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&credential{}))

	if err != nil {
		panic("grpc: dial: " + err.Error())
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02T15:04:05"))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	reply, err := client.SayHello(ctx, &proto.HelloRequest{
		Name: "xiaoming",
	})
	if err != nil {
		panic("grpc: sayHello: " + err.Error())
	}

	fmt.Println(reply.Message)
}
