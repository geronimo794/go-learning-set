package main

import (
	"fmt"
	"log"
	"os"

	"github.com/geronimo794/go-learning-set/package01basic"
)

/**
* GOLANG PACKAGE
**/
// To use golang package feature, first you need to initialize your application with
// command : go mod init (NAME_OF_YOUR_APPLICATION)
// example #1 : go mod init my-application
// example #2 : go mod init github.com/geronimo794/go-learning-set
// Example number 2 you can use if you want to publice your repository

// After that, you can create a folder to have a different package name
// example #1 : go mod init my-application/another-package
// example #2 : go mod init github.com/geronimo794/go-learning-set/package01basic
// Voila, now you can use package or module for your application
func main01CallPackage() {
	// To call the package function first you need type the package name, and then the property or the function name
	fmt.Println(package01basic.CalculateNumberPub(1, 2, 3, 4))

	// We just can call the property of package with upper case first letter
	// fmt.Println(package01basic.calculateNumberPriv(1, 2, 3, 4)) // <-- This will throw error
	// package01basic have calculateNumberPriv function,
	// but we can't call it because it's private indicated from first letter of function or property

	// Call another public function with initialize
	fmt.Println(package01basic.GetName())
}

// Instead of create our package or module we can use
// Package created by golang by default
// OS Package : Package to interact with your operationg system operation
func main01PackageOs() {
	/**
	* OS File
	**/

	// Open file in your file system
	file, err := os.Open("README.MD") // File will return *os.file poimting to the file in the memory
	if err != nil {
		log.Fatal(err)
	}

	// To read the file we need make variable the data inside the content
	// We need slice of byte
	var data []byte
	// After we define the slice of byte,
	// We need to define the length of slice of byte
	data = []byte{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // Define with how many byte you want to keep
	}
	fmt.Println("Data before read file", data)
	// After define the length of length, we will put the data to slice of byte by length of slice
	count, err := file.Read(data) // Put slice of byte, and write the slice of byte variable with the data
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %q\n", count, data[:count])
	fmt.Println("Data after read file", data)

	///////////////////////////////////////////////////////////////////

	// We can simplify the declaration of variable slice with make function
	data = make([]byte, 100) // Instead of manual loop to define the length

	fmt.Println("Data before read file", data)
	// After define the length of length, we will put the data to slice of byte by length of slice
	count, err = file.Read(data) // Put slice of byte, and write the slice of byte variable with the data
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %q\n", count, data[:count])
	fmt.Println("Data after read file", data)

	/**
	* OS Args
	**/
	// With os package we can get argument from the terminal when we execute the application
	var varArgs = os.Args
	fmt.Println("Terminal Arguments :", varArgs) // To add argument try add argument after file name : go run main.go 01_basic2.go this is an argument

}
