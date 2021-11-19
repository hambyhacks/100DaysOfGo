package main

import "fmt"

func main() {
	// Maps are like dictionaries in Python.
	// Literal syntax goes here...
	tmInternGrades := map[string]float32{
		"Kyle":     1.0,
		"Jonathan": 2.5,
		"Cyrus":    2.5,
		"Joy":      1.0,
		"Katkat":   3.0,
	}

	// making a map using make() function

	entrepreneurs := make(map[string]bool)
	entrepreneurs = map[string]bool{
		"Joy":  true,
		"Kyle": true,
	}

	// adding values to the map
	entrepreneurs["Katkat"] = true

	fmt.Printf("Grades: %v\n", tmInternGrades)
	fmt.Printf("Entrepreneurs: %v\n", entrepreneurs)

	// Deleting values from maps

	delete(tmInternGrades, "Jonathan")
	fmt.Println(tmInternGrades)

	// comma ok variable -> used to query our map.
	interns, ok := tmInternGrades["Jonathan"]
	fmt.Println(interns, ok)

	//maps also have built-in len() function which returns the size of the map
	fmt.Printf("Length of the entrepreneurs map is: %v", len(entrepreneurs))
}
