package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	var i interface{} = "true"
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Don't Know")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{} // can use any other type which can be associated with methods

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
