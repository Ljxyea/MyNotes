package factory

import "fmt"

func Do() {
	ak47 := getGun("ak47")
	muskets := getGun("musket")
	fmt.Println(ak47.getName())
	fmt.Println(muskets.getName())
	fmt.Println(ak47.getPower())
	fmt.Println(muskets.getPower())

}
