package main

import (
	"fmt"
	"strconv"
)

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

func main01VariableTypeConversion() {
	// Even same integer variable,
	// to operate between integer variable you need to convert it to exact same variable
	var varInt int = 10
	var varInt32 int32 = 10
	var varInt64Result int64

	// You can't use it directly, because the variable is not exact same type
	// varInt64Result = varInt + varInt32 // <-- This will be throw error on compilation
	varInt64Result = int64(varInt) + int64(varInt32) // You need to convert it first

	fmt.Println(varInt64Result)

	// String variable
	var varStr string = "Rozi"
	var varStrChar byte = varStr[0]

	fmt.Println(varStrChar)         // You will print the ASCII value number of the character
	fmt.Println(string(varStrChar)) // You will print the character string

	// String to integer
	varStr = "2000" // Comment this if you want to see error message

	// var varStrToInt int = int(varStr) // You can't parse it directly to became integer, you need additional function
	varStrToInt, err := strconv.Atoi(varStr) // You can't parse it directly to became integer, you need additional function

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(varStrToInt)
	}
}

func main01TypeDeclaration() {
	// Struct of identity
	type identity struct {
		name      string
		id_number int
	}
	// Type declaration is mechanism to give variable a new name or a new alias
	type kata string
	type angka int
	type ktp identity

	// You can use like a normal variable
	var kalimat kata = "Kalimat adalah susunan kata"
	var jumlah angka = 1
	var jumlah_total angka = 1

	// Operating like the real variable type
	jumlah_total = jumlah + 1

	// Println all data
	fmt.Println(kalimat)
	fmt.Println(jumlah)
	fmt.Println(jumlah_total)
}
func main01MathematicOperator() {
	// You can use mathematical operator with all number variable
	var varFloatA float32 = 10.20
	var varFloatB float32 = 6.40
	var varFloatAdd float32 = varFloatA + varFloatB      // Add operator +
	var varFloatMin float32 = varFloatA - varFloatB      // Min operator -
	var varFloatMultiply float32 = varFloatA * varFloatB // Multiply operator *
	var varFloatDivide float32 = varFloatA / varFloatB   // Divide operator /
	// var varFloatMod float32 = varFloatA % varFloatB // Except for mod operator, you just only use it with integer

	var varIntA int = 21
	var varIntB int = 5
	var varIntMod int = varIntA % varIntB // You just can only use modulus operator with integer

	fmt.Println(varFloatAdd, varFloatMin, varFloatMultiply, varFloatDivide)

	fmt.Println(varIntMod)

	// Augmented operation : Is a short operation you can use the operate the value of variable with another value
	varIntA += 10 // Equal to : varIntA = varIntA + 10
	fmt.Println(varIntA)

	varIntA -= 10 // Equal to : varIntA = varIntA - 10
	fmt.Println(varIntA)

	varIntA *= 10 // Equal to : varIntA = varIntA * 10
	fmt.Println(varIntA)

	varIntA /= 10 // Equal to : varIntA = varIntA / 10
	fmt.Println(varIntA)

	varIntA %= 10 // Equal to : varIntA = varIntA * 10
	fmt.Println(varIntA)

	// Unary Operator :	Is a operation with single operan
	varIntA++ // Equal to : varIntA = varIntA + 1
	fmt.Println(varIntA)

	varIntA-- // Equal to : varIntA = varIntA - 1
	fmt.Println(varIntA)

	varIntA = -varIntA // Equal to : varIntA = varIntA * -1
	fmt.Println( varIntA)

	varIntA = +varIntA // Equal to : varIntA = varIntA * +1
	fmt.Println(varIntA)

	var varBool bool = true
	varBool = !varBool // Equal to : If value is true it will be false, if value is false it will be true
}

func main01ComparisonOperator() {
	// You can compare number string variable
	
}
