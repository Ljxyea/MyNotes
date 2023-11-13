package decorator

import "fmt"

// 装饰模式
func Do() {
	pizza := &VeggieMania{}

	pizzaWithChinese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithChineseAndTomato := &TomatoTopping{
		pizza: pizzaWithChinese,
	}

	fmt.Printf("tomaot %d", pizzaWithChineseAndTomato.getPrice())
}
