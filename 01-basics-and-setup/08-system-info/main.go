package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("System OS name: ", runtime.GOOS)
	fmt.Println("System Architecture: ", runtime.GOARCH)
	fmt.Println("System number of CPUs: ", runtime.NumCPU())

	// output
	// 	System OS name:  windows
	// System Architecture:  amd64
	// System number of CPUs:  12
}
