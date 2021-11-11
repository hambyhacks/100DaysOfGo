package main

import "fmt"

// Package Scope Declarations
var I, J float32 = 77.7, 66.6 // Notice the uppercase variables "I" and "J", we also used them in add() function

func main() {
	var s string = "Hello Gophers!"
	var q float32 = 10.20
	var w float32 = 20.30
	fmt.Printf("The value of variable s is: %v\n", s) // variable "s" is declared @main scope
	fmt.Printf("Answer is equals to: %v", add(q, w))
}

func add(x, y float32) float32 {
	x = I + J // we passed values of variable "q" and "w" to add function
	y = I - J //but we shadowed them by adding/subtracting variables "I" and "J" and storing it into variables "x" and "y"
	ans := x * y
	return ans // ans value will be returned when called in function @main
}
