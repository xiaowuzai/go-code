package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/xiaowuzai/go-code/learn_grpc/grpc_metadata/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
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
