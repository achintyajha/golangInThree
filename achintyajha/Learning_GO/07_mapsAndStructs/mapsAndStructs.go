package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	// maps --> like dictionaries
	statePopulation := map[string]int{
		"California":   123123,
		"Texas":        74123123,
		"Florida":      23123123,
		"New York":     412324233123,
		"Pennsylvania": 3423423123,
		"Illinois":     4343423,
		"Ohio":         122424234,
	}
	statePopulation["Georgia"] = 23423423 // changes order
	delete(statePopulation, "Florida")
	fmt.Println(statePopulation["Florida"]) // Zero, problem as this means either FLorida is not in map or has population equal to zero
	fmt.Println(statePopulation)
	fmt.Println(statePopulation["Ohio"])

	// Workaround for detecting zero value or missing key
	population, ok := statePopulation["Florida"]
	fmt.Println(population, ok) // ok prints false

	// Maps change values when passed to other variables or funcs

	// structs, only collection where elements can have different data types

	aDoctor := Doctor{
		number:    3, //not necessary to use 'number:' in syntax, but it is good for readability. Also, it is positional so not writing name can cause problems
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	// Anonymous struct, use in short-lived cases
	anotherDoctor := struct{ name string }{name: "Jon Pertwee"}
	fmt.Println(anotherDoctor)

	// Structs don't change values when passed to other variables or funcs, make copies

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])
}
