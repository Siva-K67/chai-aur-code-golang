package main

import (
	"fmt"
	"os"
)

func main() {
	//stage 1- writing a file
	// writing into a a file is ultimately a byte-level operation
	// converts readable string into byte
	content := []byte("Hello from Ippo Makunochi!")
	err := os.WriteFile("output.txt", content, 0644)
	if err != nil {
		fmt.Println("error writing file: ", err)
		return
	}
	fmt.Println("file written succesfully")

	//stage 2- reading the same file
	//string(data) converts those raw bytes back into a readable Go string
	data, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}
	fmt.Println("file contents: ", string(data))
}
