package simple_factory

import "fmt"

const TypeHi = 1
const TypeHello = 2

// API 接口定义
type API interface {
	Say(name string) string
}

// NewAPI 根据类型返回 API 实例
func NewAPI(t int) API {
	switch t {
	case TypeHi:
		return &hiAPI{}
	case TypeHello:
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (h *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
