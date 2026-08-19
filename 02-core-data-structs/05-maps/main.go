package main

import "fmt"

func main() {

	// map literal - key type, value type
	//auto sorts alphabetically
	ages := map[string]int{
		"Siva":   26,
		"Sampat": 24,
		"Anshul": 24,
		"Shinde": 25,
	}
	fmt.Println(ages)

	//make-creates an empty map
	scores := make(map[string]int)
	scores["math"] = 90
	scores["scince"] = 86
	scores["english"] = 100
	scores["hindi"] = 85
	fmt.Println(scores)

	//accessing a value
	fmt.Println(ages["Shinde"])

	//comma-ok, checking if key exists
	value, exists := ages["Utkarsh"]
	fmt.Println(value, exists)

	//deleting a key
	delete(scores, "math")

	//looping over a map with range
	for key, val := range scores {
		fmt.Println(key, val)
	}

	//lenght of a map
	fmt.Println(len(ages))

	/*

		map[Anshul:24 Sampat:24 Shinde:25 Siva:26]
		map[english:100 hindi:85 math:90 scince:86]
		25
		0 false
		scince 86
		english 100
		hindi 85
		4


	*/

}
