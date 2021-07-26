package main

import (
	"fmt"
)

func main() {

	statePopulations := map[string]int{
		"California":   123123,
		"Texas":        74123123,
		"Florida":      23123123,
		"New York":     412324233123,
		"Pennsylvania": 3423423123,
		"Illinois":     4343423,
		"Ohio":         122424234,
	}

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 5

	if guess < 1 {
		fmt.Println("The guess must be more than 1.")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100.")
	} else {
		if guess < number {
			fmt.Println("Too low!")
		}
		if guess > number {
			fmt.Println("Too High!")
		}
		if number == guess {
			fmt.Println("You got it!")
		}
	}
}
