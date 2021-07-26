package main

import (
	"fmt"
	"strconv"
)

var i int = 27

var k int = 0 //unused global variable, not an error

func main() {
	// j := 13 //unused local variable, an error
	var i int = 11
	i = 22
	println(i)

	var m float32
	m = 24
	fmt.Printf("%v, %T", m, m)
	println()

	m = float32(i)
	fmt.Printf("%v, %T", m, m)
	println()

	var str string
	str = strconv.Itoa(i)
	fmt.Printf("%v, %T \n", str, str)
}

// Naming conventions
// Camel case or pascal convention
// variable used for shorter duration --- small name (example using i for counter in 'for-loop')
// variable used many times and for longer duration --- more meaningful name(example gameStatus)
// acronyms all upper case -- var goodURL string = "asdadasdasdsad"
