package bridge

import "fmt"

//桥接模式: 将功能分成抽象部分和实施部分, 独立开发

// 抽象
type Computer interface {
	Print()
	SetPrinter(p Printer)
}

// 精确抽象
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Mac Print")
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// 精确抽象
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Windows Print")
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// 实施
type Printer interface {
	PrintFile()
}

// 具体实施
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Epson ")
}

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Hp ")
}
