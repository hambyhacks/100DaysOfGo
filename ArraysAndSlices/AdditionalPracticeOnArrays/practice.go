package main

import "fmt"

func main() {
	a := make([]int,0, 50) //make zero length slice with capacity of 50
	b := []int{-10,-9,-8,-7,-6,-5,-4,-3,-2,-1} // slice of integers from -10 to -1, Len: 10, Cap:10
	c := []int{1,2,3,4,5,6,7,8,9,10} // slices of integers from 1 to 10, Len: 10, Cap:10
	
	d := append(a,b...) // appends slice b to slice a, results to: [-10,-9,-8,-7,-6,-5,-4,-3,-2,-1] len=10, cap=50
	e := append(c,a...) // appends slice b to slice c, results to: [1,2,3,4,5,6,7,8,9,10] len=20, cap=20 
	f := append(d,e...) // appends slice d to slice e, results to: [-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 1 2 3 4 5 6 7 8 9 10], len=20, cap=50
	// capacity in line 13&14 happens because slices are dynamically sized.
	fmt.Printf("Length: %v Capacity: %v\n",len(d),cap(d))
	fmt.Printf("Length: %v Capacity: %v\n",len(e),cap(e))
	fmt.Printf("Length: %v Capacity: %v\n",len(f),cap(f))


	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
