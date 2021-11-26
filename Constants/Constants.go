package main

import (
	"fmt"
)

const a int64 = 322
const (
	_  = iota // 0 and assigned as blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	EB
)
const (
	_     = iota // special symbol for related constants to be created easily
	Admin = 1 << iota
	Finance
	InfoSec
	Executive
)

func main() {
	// Naming Conventions
	const piValue float32 = 3.1416 // Do not start with uppercase letter when declaring constant if you only want to use it internally in a function.
	fmt.Printf("%v: %T\n", piValue, piValue)

	// Constants can be of any primitive data types and also can be shadowed.

	fmt.Printf("%v: %T\n", a, a) // Gets the value of const a from declaration in package scope.

	const a int = 14 // Shadowed the value of const a in package scope.
	const b string = "Go!"
	const c float64 = 2.455
	const d bool = true

	fmt.Printf("%v: %T\n%v: %T\n%v: %T\n%v: %T\n", a, a, b, b, c, c, d, d)

	// Untyped constants (can work as literals)

	const e = 35.0          // float32
	const f float64 = 52.90 //float64

	fmt.Printf("Untyped Constants can be used to work as literals! ex: e+f=%v\n", e+f)

	// Enumerated Constants through bit shifting example
	fileSize := 20000000.00
	fmt.Printf("%.2fMB\n", fileSize/MB)

	jobRole := Admin | Finance                                // Sets bit for admin and finance role
	fmt.Printf("%b\n", jobRole)                               // prints in byte form
	fmt.Printf("Infosec Admin? %v", Admin&jobRole == InfoSec) // checks if the jobRole is equal to bits of an InfoSec Admin.
}
