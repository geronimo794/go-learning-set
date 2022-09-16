package main

import "fmt"

func main01HelloWorld() {
	// Use fmt package to print text to console
	fmt.Println("Hello World")

	// You can use comma to print separated data
	fmt.Println("Hello", "Separated", "World") // Output data will be auto separated with space

	// You can print any variable with fmt.Println : fmt.Println("Hello", varA, structB, interfaceC)

	// To run your program use : go run main.go 01_basic.go

	// To generate executabel of the program use : go build main.go 01_basic.go
}

func main01VariableAndConstantDeclaration() {
	// There is 4 ways to declare variable
	// 4 : var (variable_name) (variabel_type)
	var varA string
	varA = "a"

	// 2 : var (variable_name) = (variabel_value)
	var varB = "a" // Go lang will automatically detect the variable type

	// 3 : (variable_name) := (variabel_value)
	varC := "c"

	// 4 : Multiple variable declartion
	var (
		varD string = "d"
		varE string = "e"
		varF string = "f"
	)

	// Constant is a fixed variable data. So you can't change the value
	// and you need to declare the value directly.
	const constA string = "a"
	// You can't change the value
	// constA = "X" // <-- This will be throw error

	// EVERY variable but not CONSTANT need to be use on Golang
	// If not, the compilation will be error
	fmt.Println(varA)
	fmt.Println(varB)
	fmt.Println(varC)
	fmt.Println(varD, varE, varF)
}

func main01VariableTypeNumberAndBoolean() {
	// There 2 main type of golang variable : integer and floating point
	// Integer
	var varInt int = 1                       // minimal int32 : -2147483648 until 2147483647
	var varInt8 int8 = 127                   // -128 until 127
	var varInt16 int16 = 32767               // -32768 until 32767
	var varInt32 int32 = 2147483647          // -2147483648 until 2147483647
	var varInt64 int64 = 9223372036854775807 // -9223372036854775808 until 9223372036854775807

	// Unsigned Integer : Integer that have no minus value
	var varUInt uint = 1                        // minimal uint32 : 0 until 4294967295
	var varUInt8 uint8 = 255                    // 0 until 255
	var varUInt16 uint16 = 65535                // 0 until 65535
	var varUInt32 uint32 = 4294967295           // 0 until 4294967295
	var varUInt64 uint64 = 18446744073709551615 // 0 until 18446744073709551615

	// Floating Point : number with comma
	// TBC
	var varFloat32 float32 = 1.0
	var varFloat64 float64 = 1.0
	var varComplex64 complex64 = 1.0
	var varComplex128 complex128 = 1.0

	// Type Alias, or same value but just different call name
	var varByte byte = 1 // uint8
	var varRune rune = 1 // int32

	// If your variable overflow receive more than it can hold. It will be back to start point
	// varInt8 127 ... -128
	// varUInt8 255 ... 0
	varInt32 = varInt32 + varRune
	varUInt8 = varUInt8 + varByte

	// Dont forget to use all your variable and remove the unused one
	fmt.Println(varInt, varInt8, varInt16, varInt32, varInt64)
	fmt.Println(varUInt, varUInt8, varUInt16, varUInt32, varUInt64)
	fmt.Println(varFloat32, varFloat64, varComplex64, varComplex128)

	// Boolean variabel that can hold value true or false
	var varBoolTrue bool = true // With default declaration
	varBoolFalse := false       // With simple short declaration
	fmt.Println(varBoolTrue, varBoolFalse)
}

func main01VariableTypeString() {
	// A variable string on golang, you need to use double quote
	// var varStr string = 'This is string' // <-- This is will be failed
	var varStr string = "This is string"

	// String operation
	// 1. len : Count the length of the string
	varLenStr := len(varStr) // The return value will be int

	// 2. varStr[x] : Get the character on position
	varCharStr := varStr[0] // The return value will be byte(uint8)

	fmt.Println(varStr)
	fmt.Println(varLenStr, varCharStr)
}
