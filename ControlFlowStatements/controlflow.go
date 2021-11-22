package main

import (
	"fmt"
)

func main() {
	// we can use , ok statements to query a map and check in an if conditional block.
	breedsofCats := map[string]int{
		"Maine Coon":  1,
		"Orange":      3,
		"Turtleshell": 2,
		"Calico":      5,
		"Tilapia":     10,
	}
	if cats, ok := breedsofCats["Tilapia"]; ok {
		fmt.Printf("Number of Tilapia Cats: %v\nAre they ok?: %v\n", cats, ok)
	}
	// Comparison operators
	var numberTest string
	fmt.Print("Enter 11-Digit number: ")
	fmt.Scanf("%s", &numberTest)

	// Checking if length of user input is exactly 11 characters.
	if len(numberTest) == 11 {
		fmt.Printf("Contact No: %v\n", numberTest)
	} else if len(numberTest) < 11 {
		fmt.Println("Contact Number you inputted is less than 11.")
	} else {
		fmt.Println("Contact Number you inputted is greater than 11.")
	}

	// Switch Statments using tag syntax
	var pref string
	fmt.Print("Which do you prefer? Use numbers to pick.\n1. Golang\n2. Python\nYour Choice: ")
	fmt.Scanf("%s", &pref)

	if len(pref) == 1 {
		switch pref {
		case "1":
			fmt.Println("User preferred Golang.")
		case "2":
			fmt.Println("User preferred Python. Gopher will be sad!")
		default:
			fmt.Println("Select from the choices!")
		}
	} else {
		fmt.Println("Invalid choice length!")
	}

	// Switch Statements using tagless syntax and fallthrough(equivalent to JMP in assembly)
	var bookpref string
	fmt.Print("Which do you prefer? Use numbers to pick.\n1. C++\n2. Assembly\nYour Choice: ")
	fmt.Scanf("%s", &bookpref)

	if len(bookpref) == 1 {
		switch {
		case bookpref == "1":
			fmt.Println("C++ is <3")
			fallthrough
		case bookpref == "2":
			fmt.Println("but Assembly is fun :D")
		default:
			fmt.Println("Select from the choices!")
		}
	} else {
		fmt.Println("Invalid choice length!")
	}
	checkType("123")
}

// checkType Function for type switch
func checkType(a interface{}) {

	switch a.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	case float64:
		fmt.Println("Type: float64")
	default:
		fmt.Println("unknown type")
	}
}
