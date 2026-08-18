package main

import (
	"fmt"
	"time"
)

func main() {

	//current time
	current_time := time.Now()
	fmt.Println(current_time)

	//individual components
	fmt.Println(current_time.Year(), current_time.Month(), current_time.Day())
	fmt.Println(current_time.Hour(), current_time.Minute(), current_time.Second())

	//output
	// 2026-08-18 16:14:54.39867 +0530 IST m=+0.000000001
	// 2026 August 18
	// 16 14 54

}
