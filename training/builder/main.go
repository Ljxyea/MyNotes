package builder

import "fmt"

func Do() {
	normalBuilder := NewBuilder("normal")
	iglooBuilder := NewBuilder("igloo")

	director := newDirector(normalBuilder)
	nomalHouse := director.buildHouse()

	director = newDirector(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("normal house: %+v \n", nomalHouse)
	fmt.Printf("igloo house: %+v \n", iglooHouse)
}
