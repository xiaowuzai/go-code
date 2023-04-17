package main

import (
	"github.com/xiaowuzai/go-code/learn_consul/register"
)

func main() {
	reg := register.Registry{
		Host: "127.0.0.1", // consul host
		Port: 8500,        // consul http port
	}
	_ = reg.Register("192.168.204.82", 8003, "bond-server", []string{"service", "bond"}, "bond-server")
}
