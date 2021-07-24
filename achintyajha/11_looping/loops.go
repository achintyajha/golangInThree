package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("================================")
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// while - like loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//infinite loop
	k := 0
	for {
		println("================================")
		k++
		if k == 5 {
			break // breaking the infinite loop
		}
	}
	println("********************************")
	// Labels
Loop:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if i*j == 3 {
				// break; // breaks nearest loop only
				// To break outer loop --->
				break Loop
			}
		}
	}

	// Iterating overs collections
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

}

// i++ is a statement not an expression (in GO)
