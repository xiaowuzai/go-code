package register

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func NewRegistry(host string, port int) (RegistryClient, error) {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", host, port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Registry{client}, nil
}

//Register 注册服务
func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC: fmt.Sprintf("%s:%d", address, port),
		// HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err := r.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

//DeRegister 取消注册
func (r *Registry) DeRegister(serviceId string) error {
	return r.Agent().ServiceDeregister(serviceId)
}

// PrintServices 打印注册的服务
func (r *Registry) PrintServices() error {
	m, err := r.Agent().Services()
	if err != nil {
		return err
	}
	for k, v := range m {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}
	return nil
}
