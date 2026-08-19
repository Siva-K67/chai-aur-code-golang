package main

import "fmt"

//structure- a custom type data structure
type Student struct {
	Name  string
	Age   int
	Marks float64
}

func main() {

	// creating a struct with field names
	s1 := Student{
		Name:  "Utkarsh",
		Age:   26,
		Marks: 99.86,
	}

	fmt.Println(s1)

	//accessing fields with dot notation
	fmt.Println(s1.Name, s1.Age)

	//modifying any field
	s1.Marks = 98.23
	fmt.Println(s1)

	// creating without fieeld names, order matters
	s2 := Student{"Anshul", 24, 95.34}
	fmt.Println(s2)

	// all fields get their type's zero value
	var s3 Student
	fmt.Println(s3)

	//nested struct. no inheritance in golang. functions
	// cant be declared inside struct. so we declare them outside
	type Address struct {
		City  string
		State string
	}

	type Employee struct {
		Name    string
		Address Address
	}

	e1 := Employee{
		Name:    "Purohit",
		Address: Address{City: "Bhopal", State: "MAdhya Pradesh"},
	}
	fmt.Println(e1)

}
