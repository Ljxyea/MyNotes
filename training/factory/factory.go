package factory

// 产品接口
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// 具体产品
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// 具体产品
type AK47 struct {
	Gun
}

func newAK47() *AK47 {
	return &AK47{
		Gun: Gun{
			name:  "AK47",
			power: 50,
		},
	}
}

// 具体产品
type musket struct {
	Gun
}

func newMusket() *musket {
	return &musket{
		Gun: Gun{
			name:  "Musket",
			power: 30,
		},
	}
}

// 工厂
func getGun(gunType string) IGun {
	if gunType == "AK47" {
		return newAK47()
	}
	if gunType == "Musket" {
		return newMusket()
	}
	return nil
}
