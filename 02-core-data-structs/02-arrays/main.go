package main

import "fmt"

func main() {

	//declared with fixed size
	var scores [5]int
	fmt.Println(scores)

	// assign vals by index
	scores[0] = 90
	scores[1] = 93
	fmt.Println(scores)

	//arr literal- size and val together
	names := [3]string{"Endou Mamorou", "Gounji Shuuya", "Fubuki Shirou"}
	fmt.Println(names)

	// let Go count the size
	nums := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums, len(nums))

	//looping with range
	for i, name := range names {
		fmt.Println(i, name)
	}

	//arr are copied by value not by reference
	// it means that changing copyArr does not change
	// the original at all
	original := [3]int{6, 3, 4}
	copyArr := original
	copyArr[0] = 777
	fmt.Println(original, copyArr)

	/*
		[0 0 0 0 0]
		[90 93 0 0 0]
		[Endou Mamorou Gounji Shuuya Fubuki Shirou]
		[1 2 3 4 5 6 7] 7
		0 Endou Mamorou
		1 Gounji Shuuya
		2 Fubuki Shirou
		[6 3 4] [777 3 4]
	*/
}
