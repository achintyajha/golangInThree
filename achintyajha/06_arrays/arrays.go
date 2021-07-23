package main

import (
	"fmt"
)

func main() {
	grades := [...]int{97, 98, 100}
	fmt.Printf("Grades: %v\n", grades)

	students := [3]string{}
	fmt.Printf("students: %v\n", students)
	students[0] = "Achintya"
	students[1] = "Atharva"
	students[2] = "Arnav"
	fmt.Printf("students: %v\n", students)
	fmt.Printf("students #1: %v\n", students[0])
	fmt.Printf("Number of students: %v\n", len(students))
	fmt.Printf("================================================================\n")
	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)
	// Arrays and slices
	// Array ==> var array [...] = [3]int{}
	// Slice ==> var array [] = [3]int{}
	// Assigning arrays to vars created copy, but for slices it points to same location

	// make
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("capacity: %v\n", cap(a))

	// a = append(a, []int{1, 2, 3, 4, 5, 6, 7}) // error
	a = append(a, []int{1, 2, 3, 4, 5, 6, 7}...) // splits array into integers
	fmt.Println(a)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("capacity: %v\n", cap(a))
}
