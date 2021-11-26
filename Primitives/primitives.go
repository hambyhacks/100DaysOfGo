package main

import (
	"fmt"
)

func main() {
	var x bool // Golang initialize variables to 0 (false in bool)
	var y float32 = 32.213123712
	fmt.Printf("%v %T\n", y, y)
	z := int(y) // float32 to int conversion
	fmt.Printf("The variable z is converted to integer which has value: %v and type: %T\n", z, z)
	fmt.Printf("%v %T\n", x, x)

	//Bitwise operation

	a := 15 // 1111
	b := 10 // 1010

	fmt.Println(a & b)  // AND operator, if both bits are not set to 1(true), then its false. Answer: 1010 = 10
	fmt.Println(a | b)  // OR operator, if one of the bits are set to 1, then its true. Answer: 1111 = 15
	fmt.Println(a ^ b)  // XOR operator, if both of the bits are either both 0 and 1, then its false. Answer: 0101 = 5
	fmt.Println(a &^ b) // equivalent to (a & ~b), read as "a AND NOT B", equate b first for NOT operator then AND it on a. Answer: 0101 = 5

	// Bit Shifting
	// Explanation: basically move all bits according to what direction you indicated either left (<<) or right (>>)
	c := 15 // 1111

	fmt.Println(c << 8) // Shift 8-bits to the left. Answer: 111100000000 = 3840
	fmt.Println(c >> 8) // Shift 8-bits to the right.

	// Floating-Point literals & Arithmetic Operations

	d := 3.1416       // simplified pi
	e := 6.626068e-34 // Planck's Constant
	f := 10e32        // Planck Temperature

	fmt.Printf("Value: %v, Type: %T\n", d, d)
	fmt.Printf("Value: %v, Type: %T\n", e, e)
	fmt.Printf("Value: %v, Type: %T", f, f)

	fmt.Println(d + e + f)
	fmt.Println(d - e - f)
	fmt.Println(d * e * f)
	fmt.Println(d / e / f)

	// Complex Data Type & Arithmetic Operations & Getting the Real and Imaginary part of a complex number

	g := 3 - 9i
	h := 7 + 2i

	fmt.Println(g + h)
	fmt.Println(g - h)
	fmt.Println(g * h)
	fmt.Println(g / h)
	fmt.Printf("%v %v %v %v\n", real(g), imag(g), real(h), imag(h))

	// making complex number
	var i complex128 = complex(8, -12)
	fmt.Printf("%v %T\n", i, i)

	// String (double quotes to declare)

	j := "Hello, I am a string!"
	k := "Hello, Gophers!"
	fmt.Printf("%v: %T\n", j, j)
	fmt.Printf("%v: %T\n", string(j[1]), j[1]) // j[2] is a byte and byte is an alias of string so we can convert it back to string.

	// String Concatenation ("adding" strings)
	fmt.Println(j, " "+k) //adding space between concatenation

	// Slices (collection of bytes)
	l := []byte(k)
	fmt.Println(l) // Hello, Gophers string represented in decimal form, see https://www.alpharithms.com/ascii-table-512119/
	fmt.Printf("UTF-8(uint8) characters: %v, converted to string: %v\n", l, string(l))

	// Rune (UTF-32) (single quotes to declare) (Type alias to int32)
	m := 'I'
	fmt.Println(m)

}
