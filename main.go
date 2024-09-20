package main

import (
	"fmt"

	"main.go/intf"
)

func main() {
	duck := intf.NewDuck()
	john := intf.NewNamedDuck("john")
	farmer := intf.NewFarmer(duck, john)
	fmt.Println(farmer.Breed())
}
