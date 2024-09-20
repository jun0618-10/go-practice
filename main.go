package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(3, 4))

	// var i, j int = 1, 2
	// k := 3
	// fmt.Println(i, j, k)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}

func add(a, b int) int {
	return a + b
}
