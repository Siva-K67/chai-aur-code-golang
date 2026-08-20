package main

import "fmt"

func main() {
	//for loop, Go's only loop keyword
	for i := 1; i < 5; i++ {
		fmt.Println("i: ", i)
	}

	//while loop
	j := 0
	for j < 3 {
		fmt.Println("j: ", j)
		j++
	}

	//infinite loop with break
	k := 0
	for {
		if k == 2 {
			break
		}
		fmt.Println("k: ", k)
		k++
	}

	//continue
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd :", i)
	}

	//nested loop with labeled break
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer
			}
		}
		fmt.Println("nested: ", i, j)
	}

	//goto- jumps to a labeled line
	n := 0
loop:
	if n < 3 {
		fmt.Println("Goto n: ", n)
		n++
		goto loop
	}

	// use of _
	fruits := []string{"apple", "banana", "orange"}

	//print with index
	//range gives 2 values, index and value itself at that index
	fmt.Println("----with index----")
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	//we dont want the index, so we just replace it by _
	fmt.Println("----Without index----")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// what happens if you declare i but never use it (uncomment to see the error)
	// for i, fruit := range fruits {
	// 	fmt.Println(fruit)
	// }
	// ^ this would fail to compile: "i declared and not used"

}
