package main

import "fmt"

func main() {

	day := 3

	//Go needs no break statement unlike C++
	switch day {

	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("WEdnestday")
	case 4:
		fmt.Println("thursday")
	case 5, 6, 7:
		//note here mutiple values in one case
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown day or Wrong entry")

	}
}
