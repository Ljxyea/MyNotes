package main

import (
	"training/abstract_factory"
	"training/adapter"
	"training/bridge"
	"training/builder"
	"training/composite"
	"training/decorator"
	"training/factory"
)

func main() {
	//创建型
	//抽象工厂
	abstract_factory.Do()
	//工厂方法
	factory.Do()
	//生成器
	builder.Do()

	//结构型
	//适配器
	adapter.Do()
	//桥接模式
	bridge.Do()
	//组合模式
	composite.Do()
	//装饰模式
	decorator.Do()
}
