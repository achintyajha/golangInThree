package main

import (
	"fmt"
)

const (
	a = iota // iota is a counter here
	b = iota // not necessary to write iota here
	c = iota // not necessary to write iota here
)

func main() {
	// const myConstant float64 = math.Sin(1.57) // wrong as no execution can be assigned
	const myConstant int = 1 //naming convention, camel casing for constant not to be exported
	// myConstant = 27 // wrong
	fmt.Printf("%v, %T\n", myConstant, myConstant)

	//enumerated constants

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
