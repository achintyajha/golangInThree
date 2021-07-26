package main

import (
	"fmt"
)

// var actorName string = "ABC"
// var companionName string = "XYZ"
// var doctorNumber int = 3
// var season int = 11

var (
	actorName     string = "ABC"
	companionName string = "XYZ"
	doctorNumber  int    = 3
	season        int    = 11
)

var (
	counter int = 0
)

func main() {
	// var k int // initialize and declare
	// k = 22
	// fmt.Println(k)

	// var i float32 = 44 // when go doesn't have information about data type
	// fmt.Printf("%v %T", i, i)

	// fmt.Println()
	// dynamic := 33  // this is a declaration
	// fmt.Printf("%v %T", dynamic, dynamic)

	fmt.Printf("%v %T", counter, counter)
}
