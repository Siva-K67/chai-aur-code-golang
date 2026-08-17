package main

import "fmt"

func main() {

	//var with type mentioned by us
	var name string = "Siva"

	//var golanfg itself inferes
	var age = 23

	//short declaration
	city := "KAtamandu"

	//values cant be altered cof consts
	const pi = 3.14159

	fmt.Println(name, age, city, pi)

	//multiple vars at once
	var x, y int = 1, 2
	fmt.Println(x, y)

	//declared ut not assigned
	var score int
	fmt.Println(score)
}
