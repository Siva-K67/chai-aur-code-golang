package main

import (
	"errors"
	"fmt"
)

func main() {
	// error is just a built-in TYPE in Go - it represents "something went wrong"
	var err error

	// right now err is nil - meaning "no error"
	fmt.Println(err)
	fmt.Println(err == nil)

	// errors.New creates an actual error value, with a message
	err = errors.New("something went wrong")
	fmt.Println(err)
	fmt.Println(err == nil)

	/*
		Output
		<nil>
		true
		something went wrong
		false
	*/
}
