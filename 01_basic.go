package main

import "fmt"

func main01() {
	// Use fmt package to print text to console
	fmt.Println("Hello World")

	// You can use comma to print separated data
	fmt.Println("Hello", "Separated", "World") // Output data will be auto separated with space

	// You can print any variable with fmt.Println : fmt.Println("Hello", varA, structB, interfaceC)

	// To run your program use : go run main.go 01_basic.go

	// To generate executabel of the program use : go build main.go 01_basic.go

}
