package abstract_factory

import "errors"

// 抽象工厂接口
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func NewSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	}
	return nil, errors.New("unknown brand")
}

// 具体工厂
type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoes{
		Shoe: Shoe{
			logo: "adidas logo",
			size: 10,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirts{
		Shirt: Shirt{
			logo: "adidas logo",
			size: 10,
		},
	}
}

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirts{}
}

//抽象产品

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

// 具体产品
type AdidasShoes struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

// 抽象产品
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

// 具体产品
type AdidasShirts struct {
	Shirt
}

type NikeShirts struct {
	Shirt
}
