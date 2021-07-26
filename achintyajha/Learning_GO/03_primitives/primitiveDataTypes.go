package main

import (
	"fmt"
)

func main() {
	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	var a int = 8
	var b byte = 2
	res := a + int(b)
	fmt.Printf("a plus b is: %v\n", res)

	// BitWise Operators (only for integers)
	n1 := 10          // 1010
	n2 := 3           // 0011
	println(n1 & n2)  // 0010 = 2
	println(n1 | n2)  // 1011 = 11
	println(n1 ^ n2)  // Exclusive Or -- 1001 = 9
	println(n1 &^ n2) // And Not -- 0100 = 8

	a = 8           // 2 to power 3
	println(a >> 3) // Right Bitshift 3 places -- 2^3 * 2^3 = 2^6
	println(a << 3) // Left Bitshift 3 places -- 2^3 / 2^3 = 2^0

	println("================================================================")

	n3 := 3.14 // or var n3 float64 = 3.14
	n3 = 13.7e72
	n3 = 2.1e14
	fmt.Printf("%v, %T\n", n3, n3)

	// complex numbers

	var comp complex128 = 1 + 2i // or var comp complex128 = complex(1, 2)
	fmt.Printf("%v, %T\n", comp, comp)
	fmt.Printf("%v, %T\n", real(comp), imag(comp))

	// strings and runes are also there
}
