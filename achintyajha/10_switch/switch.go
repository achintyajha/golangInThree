package main

import (
	"fmt"
)

func main() {
	switch 2 {
	// test cases in "case" must be unique
	case 1:
		fmt.Println("one")

	case 2: // can add multiple nums here
		fmt.Println("two")

	default:
		fmt.Println("not one or two")
	}

	// tagless syntax
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to 10")
		// since there is no need for break statement, there is no auto fallthrough either
		fallthrough // it can be achieved by using 'fallthrough'
		// even if next case is false, fallthrough will execute it
	case i <= 20:
		fmt.Println("Less than or equal to 20")
	default:
		fmt.Println("Greater than 20")
	}
}

// most languages have implicit fallthrough and explicit break
// go has implicit break and explicit fallthrough

// we can break out earlier as well with the break statement
