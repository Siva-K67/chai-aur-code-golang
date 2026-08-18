package main

import "fmt"

func main() {

	//convert int into float64
	var i int = 10
	var f float64 = float64(i)
	fmt.Println(f)

	//float to int ( turncates no round off)
	var g float64 = 8.92 // vit cgpa
	var j int = int(g)   // gives o/p as 8
	fmt.Println(j)

	//int to string
	// converts to ASCII Char
	//rune is basically int32 undeer the hood
	var n int = 65
	var m string = string(rune(n))
	fmt.Println(m)

	//number to string, the correct way
	// fmt.Sprintf("%d", num)
	// it takes the format string, plugs 42 into the %d slot,
	// and builds the string "42" — then returns it
	var num int = 42
	var text string = fmt.Sprintf("%d", num)
	fmt.Printf("Type of text: %T ", text)

}
