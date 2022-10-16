package strategy

// IStrategy 策略接口定义
type IStrategy interface {
	do(int, int) int
}

// 策略实现：加
type add struct{}

func (a *add) do(x, y int) int {
	return x + y
}

// 策略实现：减
type reduce struct{}

func (r *reduce) do(x, y int) int {
	return x - y
}

// Operator 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (o *Operator) setStrategy(strategy IStrategy) {
	o.strategy = strategy
}

// 调用策略中的方法
func (o *Operator) calculate(x, y int) int {
	return o.strategy.do(x, y)
}
