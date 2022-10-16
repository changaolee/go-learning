package abstract_factory

import "fmt"

const TypeHi = 1

// RestAPI 接口定义
type RestAPI interface {
	Say(name string) string
}

// RpcAPI 接口定义
type RpcAPI interface {
	Say(name string) string
}

// APIFactory 工厂方法接口定义
type APIFactory interface {
	CreateRestAPI() RestAPI
	CreateRpcAPI() RpcAPI
}

// NewAPIFactory 用一个简单工厂封装工厂方法
func NewAPIFactory(t int) APIFactory {
	switch t {
	case TypeHi:
		return &hiAPIFactory{}
	}
	return nil
}

type hiAPIFactory struct{}

func (h *hiAPIFactory) CreateRestAPI() RestAPI {
	return &hiRestAPI{}
}

func (h *hiAPIFactory) CreateRpcAPI() RpcAPI {
	return &hiRpcAPI{}
}

type hiRestAPI struct{}

func (h *hiRestAPI) Say(name string) string {
	return fmt.Sprintf("[REST] Hi, %s", name)
}

type hiRpcAPI struct{}

func (h *hiRpcAPI) Say(name string) string {
	return fmt.Sprintf("[RPC] Hi, %s", name)
}
