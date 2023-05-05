package main

import (
	"github.com/xiaowuzai/go-code/learn_consul/register"
)

func main() {
	reg, err := register.NewRegistry("127.0.0.1", 8500)
	if err != nil {
		panic(err)
	}
	// _ = reg.Register("192.168.204.82", 8088, "bond-server", []string{"service", "bond"}, "bond-server")
	_ = reg.Register("192.168.204.82", 8088, "grpc-server", []string{"service", "grpc"}, "grpc-server")

	_ = reg.PrintServices()
}
