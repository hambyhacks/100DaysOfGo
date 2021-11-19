package main

import (
	"fmt"
	"reflect"
)

type Students struct {
	id       int
	name     string
	subjects []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal // embeds Animal struct to Bird struct
	Speed  float32
	CanFly bool
}

type Books struct {
	Name string `required max: "50"`
	ID   int    `required max: "8"`
}

func main() {
	// Explicitly listing the names of fields in a struct is not required.
	student1 := Students{
		id:   1,
		name: "Katkat",
		subjects: []string{
			"Playing",
			"Eating",
		},
	}
	fmt.Println(student1.subjects)

	// Anonymous Structs

	student2 := struct{ name string }{name: "Maya"}
	fmt.Println(student2)

	student3 := struct{ id int }{}
	anotherstudent := student3
	anotherstudent.id = 99
	fmt.Println(anotherstudent)

	//Embedding using struct
	b := Bird{} // Instantiate Bird struct
	b.Name = "Eagle"
	b.Origin = "Philippines"
	b.Speed = 30.12831823718412
	b.CanFly = true
	fmt.Println(b)

	// Using tags in struct
	t := reflect.TypeOf(Books{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) // gets the tag on field named "Name"

	// Practice
	c := Students{}
	c.name = "Kyle"
	c.id = 0
	c.subjects = []string{"RE", "Active Directory Pentesting"}

	fmt.Printf("Initial Values of Students struct:%v\nPresent Value of Students struct: %v", Students{}, c)
}
