package main

import (
	"fmt"
)

func main() {
	// Arrays
	// Built-in functions for both slices and arrays: len() and cap(), length and capacity respectively.
	num := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Listing array values: %v\n", num)
	letters := [...]string{"Hello Gophers!"}
	fmt.Printf("String array: %v\n", letters)

	// Slices (denoted by square brackets [])

	pie := []string{"Moon Pie", "Custard Pie"}
	fmt.Printf("Available pies are: %v\n", pie)
	num1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// syntax: [x:y] -> x is initial element, y is "up to but not including"
	fmt.Println(num1[:])   // print all elements of the slice
	fmt.Println(num1[1:6]) // print 2nd element up to but not including 6th element

	// make function, syntax: make(slices/array, len, cap)
	a := make([]int, 5)
	fmt.Println(a)
	b := make([]int, 5, 25)
	fmt.Printf("Length: %v, Capacity: %v\n", len(b), cap(b))

	// append() function

	c := []int{}
	c = append(c, 2, 3, 4, 5)
	fmt.Println(c)
	fmt.Printf("Length: %v, Capacity: %v\n", len(c), cap(c))
}
