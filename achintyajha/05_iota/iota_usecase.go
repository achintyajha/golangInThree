package main

import (
	"fmt"
)

const (
	// '_' as a constant name is a read only constant name
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 1 times 2 to the power 10
	MB                    // 1 times 2 to the power 100
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
	println(KB)
	println(MB)
	println(GB)
	println(TB)
	println(PB)
	println(EB)
	// println(ZB) // OVERFLOWS INT64
	// println(YB)
}
