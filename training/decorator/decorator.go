package decorator

// 装饰模式
// 零件接口
type IPizza interface {
	getPrice() int
}

// 具体零件
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 10
}

// 具体装饰
type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 5
}

// 具体装饰2
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 5
}
