package main

import (
	"fmt"
)

func main() {
	a := 20
	b := a
	a = 45
	fmt.Println(a, b) // B is assigned value of a, so both are separated

	// Pointers
	var n int = 42
	var m *int = &n // address operator,  * declares that m is a pointer
	fmt.Println(&n, m)
	n = 33
	fmt.Println(n, *m) // Here, * is the dereferencing operator

	*m = 23
	fmt.Println(n, *m)
}
