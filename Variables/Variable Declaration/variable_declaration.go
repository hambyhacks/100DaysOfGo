package main

import "fmt"

func main() {
	// First Declaration
	var i int
	i = 50

	// Second Declaration
	var j float32 = 42.5

	// Third Declaration
	k := "hello"

	fmt.Printf("%v: %T\n", i, i)
	fmt.Printf("%v: %T\n", j, j)
	fmt.Printf("%v: %T\n", k, k)
}
