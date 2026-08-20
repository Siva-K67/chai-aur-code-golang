package main

import "fmt"

func main() {

	// defer is used to cleanup. like closing db, closing file after reading it etc etc
	fmt.Println("Start")

	//runs at the end of the surrounding function, not immediartely
	defer fmt.Println("Runs at the last 1")
	defer fmt.Println("Runs at the last 2")
	defer fmt.Println("Runs at the last 3")

	//defer calls run in LIFO order

	fmt.Println("endcd ")
}
