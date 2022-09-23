package main

import (
	"fmt"

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
}
