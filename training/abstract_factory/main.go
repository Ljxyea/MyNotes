package abstract_factory

import "fmt"

// 工程
func Do() {
	adidasFactory, _ := NewSportsFactory("adidas")
	nikeFactory, _ := NewSportsFactory("nike")

	nikeShoes := nikeFactory.makeShoe()
	adidasShoes := adidasFactory.makeShoe()

	nikeShirts := nikeFactory.makeShirt()
	adidasShirts := adidasFactory.makeShirt()

	printShoeDetails(nikeShoes)
	printShoeDetails(adidasShoes)

	printShirtDetails(nikeShirts)
	printShirtDetails(adidasShirts)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Shoe logo: %s\n", s.getLogo())
	fmt.Printf("Shoe size: %d\n", s.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Shirt logo: %s\n", s.getLogo())
	fmt.Printf("Shirt size: %d\n", s.getSize())
}
