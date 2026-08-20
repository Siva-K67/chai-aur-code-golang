package main

import "fmt"

//takes no parameter
func greet() {
	fmt.Println("Inazuma 11!")
}

//takes one parameters
func greetPerson(name string) {
	fmt.Println("Hi ", name)
	// this is pass by value by defalut
}

//multiple para
func addNum(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

//same type para can share same type declaration
func addNumShort(a, b int) {
	fmt.Println("Sum short: ", a+b)
}

//return a single val, type declared after the para
func add(a int, b int) int {
	return a + b
}

//return multiple values. VEry common in Golang
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

//named return values- declared in the signature, returned auto
func rectArea(length, width int) (area int) {
	area = length * width
	return
}

// varidic parameter (ccepts 0 or any nukmber of args)
func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}

	return total
}

func main() {
	greet()
	greetPerson("Shiva")
	addNum(4, 8)
	addNumShort(2, 9)

	result := add(5, 9)
	fmt.Println("return type: ", result)

	q, r := divide(6, 4)
	fmt.Println("q: ", q, "r : ", r)

	ar := rectArea(6, 7)
	fmt.Println("area of rect is: ", ar)

	fmt.Println("sum: ", sumAll(1, 2, 3, 4))
	fmt.Println("sum: ", sumAll(12, 45, 63, 22, 11, 66, 78, 21, 0, 3))
}
