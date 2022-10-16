package factory_method

import "fmt"

const TypeHi = 1
const TypeHello = 2

// API 接口定义
type API interface {
	Say(name string) string
}

// APIFactory 工厂方法接口定义
type APIFactory interface {
	CreateAPI() API
}

// NewAPIFactory 用一个简单工厂封装工厂方法
func NewAPIFactory(t int) APIFactory {
	switch t {
	case TypeHi:
		return &hiAPIFactory{}
	case TypeHello:
		return &helloAPIFactory{}
	}
	return nil
}

type hiAPIFactory struct{}

func (h *hiAPIFactory) CreateAPI() API {
	return &hiAPI{}
}

type hiAPI struct{}

func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPIFactory struct{}

func (h *helloAPIFactory) CreateAPI() API {
	return &helloAPI{}
}

type helloAPI struct{}

func (h *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
