package main

import (
	"fmt"
)

func main() {
	sayMessage("Hello", "GO")
	res := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The sum is: ", *res)

	// Anonymous functions
	func() {
		fmt.Println("Hello Go!")
	}() // invoke this functions
}

// sayMessage(greeting string, name string) // syntactic sugar below
func sayMessage(greeting, name string) {
	fmt.Println(greeting, name)

}

func sum(values ...int) *int { // slice is passed in, variadic parameters
	fmt.Println(values)
	result := 0
	for _, value := range values {
		result += value
	}
	return &result
}

// -----Not writing about methods here-----
