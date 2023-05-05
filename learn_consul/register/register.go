package register

import "github.com/hashicorp/consul/api"

type Registry struct {
	*api.Client
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
	PrintServices() error
}
