package main

import "fmt"

var i int = 5000 // declares variable "i" with value 5000 @ package scope

func main() {
	fmt.Println(i) // package scope variable "i" was used
	i := 9000      // declares a variable "i" with value 9000 @ main()
	fmt.Println(i)
}
