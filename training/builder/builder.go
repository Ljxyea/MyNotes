package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func NewBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

type House struct {
	doorType   string
	windowType string
	floor      int
}

// 具体生成器
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "木窗"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "木门"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 1
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// 具体生成器
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "玻璃窗"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "玻璃门"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 2
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// 主管
type Director struct {
	builder IBuilder
}

func newDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setWindowType()
	d.builder.setNumFloor()
	d.builder.setDoorType()
	return d.builder.getHouse()
}
