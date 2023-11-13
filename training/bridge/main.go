package bridge

// 桥接模式
func Do() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	windowComputer := &Windows{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	windowComputer.SetPrinter(epsonPrinter)
	windowComputer.Print()
}
