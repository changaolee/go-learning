package template

import "fmt"

// Cooker 全部接口定义
type Cooker interface {
	fire()
	cook()
	outfire()
}

// CookMenu 可看作一个抽象类
type CookMenu struct{}

func (c *CookMenu) fire() {
	fmt.Println("开火")
}

// 交给具体的子类实现
func (c *CookMenu) cook() {
	panic("implement me")
}

func (c *CookMenu) outfire() {
	fmt.Println("关火")
}

// DoCook 封装具体步骤
func DoCook(cook Cooker) {
	cook.fire()
	cook.cook()
	cook.outfire()
}

// Potato 使用组合来模拟继承
type Potato struct {
	CookMenu
}

func (f Potato) cook() {
	fmt.Println("做土豆")
}
