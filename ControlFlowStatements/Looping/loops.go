package main

import "fmt"

func main() {
	// Simple For loop
	i := 0
	for {
		if i == 1000 {
			fmt.Println("Hello!")
			break
		} else {
			i++
		}

	}
	// For loop with break using tags
	j := 1
	k := 10
Loop2:
	for ; j != 10; j++ { // starts at 1 and increments to 10
		for ; k != 1; k-- { // starts from 10 and decrements to 1
			if j == 10 {
				break Loop2 // stops when j is equal to 10.
			}
		}
	}
	// Using range keyword
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // slice of numbers from 0-9
	for x, y := range s {
		fmt.Println(x*2, 2*y+1) // makes the index an even number and values an odd number.
	}
}
