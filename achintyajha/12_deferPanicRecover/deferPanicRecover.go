package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("First Line") // defer works in reverse order of assignment
	defer fmt.Println("Middle Line")
	defer fmt.Println("End Line")
	panic("Something went wrong") // goes after defer

	// a := "start"
	// defer fmt.Println(a)
	// a = "end"
}
