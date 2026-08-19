package main

import "fmt"

func main() {

	// slice declaration - no sizde in the brackets
	friends := []string{"Nene", "Kazama", "Masao"}
	fmt.Println(friends, len(friends))

	//append at the back, slices can grow dynamically
	//c++ equivalent of vectors
	friends = append(friends, "Bo-chan")
	fmt.Println(friends, len(friends))

	// make create a slice with 3 boxes filled out of total 5
	//here 3 is lenght, 5 is total capacity. so here, 2 boses are empty
	snacks := make([]string, 3, 5)
	fmt.Println(snacks, len(snacks), cap(snacks))

	//slicing, taking a sub section
	classmates := []string{"Shichan", "Kazama", "NEne", "Masao", "BO-chan"}
	KasukabeGang := classmates[1:3]
	fmt.Println(KasukabeGang)

	//kasukabe isnt a seperate copy, it is the same underlying
	// sub array. i.e slices share the underlyin g array
	KasukabeGang[0] = "Ai-chan"
	fmt.Println(classmates, KasukabeGang)

	//actually duplicate the data
	dst := make([]string, len(classmates))
	copy(dst, classmates)
	dst[0] = "Himawari"
	fmt.Println(classmates, dst)

}
