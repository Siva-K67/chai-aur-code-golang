package main

import "fmt"

type Student struct {
	Name  string
	Marks float64
}

//normal stand alone function
func printInfo(s Student) {
	fmt.Println(s.Name, s.Marks)
}

// a method ( a function of a struct. but golang does not allow
// functions to be declared INSIDE structs. so we declare them outside)
// METHOD - same idea, but with a receiver: (s Student)
// value receiver - gets a COPY of the struct
func (s Student) PrintInfo() {
	fmt.Println(s.Name, s.Marks)
}

// value receiver trying to modify - won't affect the original
func (s Student) AddMarksWrong(extra float64) {
	s.Marks = s.Marks + extra
}

//pointer receiver. operates on the original structure
func (s *Student) AddMarks(extra float64) {
	s.Marks = s.Marks + extra
}

func main() {

	s1 := Student{"Siva", 100}

	// calling the function - pass the struct as an argument
	printInfo(s1)

	// calling the method - called ON the struct, dot notation
	s1.PrintInfo()

	// value receiver - original stays unchanged
	s1.AddMarksWrong(10)
	fmt.Println("after AddMarksWrong:", s1.Marks)

	// pointer receiver - original actually changes
	s1.AddMarks(10)
	fmt.Println("after AddMarks:", s1.Marks)

}
