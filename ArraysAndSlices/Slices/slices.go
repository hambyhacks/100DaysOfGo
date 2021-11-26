package main

import "fmt"

func main() {

	// using slice defaults: low bound: 0 , high bound: length
	s := []int{2, 3, 5, 7, 11, 13}
	// 0,1,2,3,4,5
	fmt.Println(len(s))
	s = s[1:4] // s[1] -> s[4] : from index 1 up to but not including index 4
	fmt.Println(s) // [3,5,7]
	
	s = s[:2] // from the slice above @ line 11, we derive the new slice containing [3,5,7]. from index 0 up to but not including index 2.
	fmt.Println(s) // [3,5]  

	s = s[1:] // from slice above @ line 14: from index 1 up to the end of the slice.
	fmt.Println(s) // [5]
}
