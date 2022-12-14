package main

import (
	"errors"
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
	varIntA -= 10 // Equal to : varIntA = varIntA - 10
	varIntA *= 10 // Equal to : varIntA = varIntA * 10
	varIntA /= 10 // Equal to : varIntA = varIntA / 10
	varIntA %= 10 // Equal to : varIntA = varIntA * 10

	// Unary Operator :	Is a operation with single operan
	varIntA++          // Equal to : varIntA = varIntA + 1
	varIntA--          // Equal to : varIntA = varIntA - 1
	varIntA = -varIntA // Equal to : varIntA = varIntA * -1
	varIntA = +varIntA // Equal to : varIntA = varIntA * +1

	var varBool bool = true
	varBool = !varBool // Equal to : If value is true it will be false, if value is false it will be true
}

func main01ComparisonOperator() {
	// You can compare NUMBER and STRING variable
	// Comparation result will be BOOLEAN
	var varStrA string = "Ach"
	var varStrB string = "Rozikin"

	var varBoolEq bool = varStrA == varStrB    // Is equal ?
	var varBoolNotEq bool = varStrA != varStrB // Is not equal ?
	var varBoolGt bool = varStrA > varStrB     // Is greather than ?
	var varBoolLt bool = varStrA < varStrB     // Is less than ?
	var varBoolGtEq bool = varStrA >= varStrB  // Is grather than equal?
	var varBoolLtEq bool = varStrA <= varStrB  // Is less than equal?

	fmt.Println("varBoolEq :", varBoolEq)
	fmt.Println("varBoolNotEq :", varBoolNotEq)
	fmt.Println("varBoolGt :", varBoolGt)
	fmt.Println("varBoolLt :", varBoolLt)
	fmt.Println("varBoolGtEq :", varBoolGtEq)
	fmt.Println("varBoolLtEq :", varBoolLtEq)

	var varIntA int = 10
	var varIntB int = 20

	// You can use comparation operator directly on if else statement without save it first on variable
	if varIntA == varIntB { // Is (varIntA) equal (varIntB) ?
		fmt.Println("varIntA == varIntB")
	}
	if varIntA != varIntB { // Is (varIntA) not equal (varIntB) ?
		fmt.Println("varIntA != varIntB")
	}
	if varIntA > varIntB { // Is (varIntA) greater than (varIntB) ?
		fmt.Println("varIntA > varIntB")
	}
	if varIntA < varIntB { // Is (varIntA) less than (varIntB) ?
		fmt.Println("varIntA < varIntB")
	}
	if varIntA >= varIntB { // Is (varIntA) greater than equal (varIntB) ?
		fmt.Println("varIntA >= varIntB")
	}
	if varIntA <= varIntB { // Is (varIntA) less than equal (varIntB) ?
		fmt.Println("varIntA <= varIntB")
	}
}

func main01BooleanOperator() {
	// After you get or you set boolean value
	// You can compare that boolean variable to another boolean variable
	// The result will be boolean too
	var varBoolTrue bool = true
	var varBoolFalse bool = false

	var varAndResult bool = varBoolTrue && varBoolFalse // TRUE and FALSE = FALSE
	var varOrResult bool = varBoolTrue || varBoolFalse  // TRUE or FALSE = TRUE
	var varNotResult bool = !varBoolTrue                // NOT TRUE = FALSE

	// Just like another boolean result, you can compare it directly on if else statement
	if varBoolTrue && varBoolFalse { // TRUE and FALSE = FALSE
		fmt.Println("varBoolTrue && varBoolTrue")
	}

	if varBoolFalse || varBoolTrue { // FALSE or TRUE = TRUE
		fmt.Println("varBoolFalse || varBoolFalse")
	}

	if !varBoolFalse { // NOT FALSE = TRUE
		fmt.Println("!varBoolFalse")
	}

	fmt.Println(varAndResult, varOrResult, varNotResult)

}
func main01ArrayAndSliceVariableType() {
	// Array is a fixed size variable
	// In other hand Slice have dynamic size
	/**
	* ARRAY
	**/
	// This is ARRAY
	// With size of 3
	var varStrArr [3]string = [3]string{
		"Ach",
		"Rozikin",
	}
	// This is another declaration of ARRAY
	// With size of 2
	var varStrArrDecl = [...]string{
		"Ach",
		"Rozikin",
		"Malang",
		"Jawa Timur",
	}
	fmt.Println(varStrArr, varStrArrDecl)

	/**
	* SLICE
	**/
	// With dynamic size
	var varStrSlice []string = []string{
		"Ach",
		"Rozikin",
	}
	fmt.Println(varStrSlice)

	// Accessing array and slice
	fmt.Println(varStrArrDecl[0], varStrSlice[0])
	// Change value of array or slice
	varStrArrDecl[0] = "ACHMAD"
	varStrSlice[0] = "ACHMAD"

	// Function on array and slice
	// Get length
	var varIntSliceLen int = len(varStrSlice)
	var varIntArrayLen int = len(varStrArr)

	// Get capacity
	var varIntSliceCap int = cap(varStrSlice)
	var varIntArrayCap int = cap(varStrArr)

	// Append data to slice
	var varStrSliceAdd []string = append(varStrSlice, "Malang", "Jawa Timur")
	// You can't append array to slic
	// var varStrStrAdd []string = append(varStrArr, "Malang", "Jawa Timur") // <-- This code will be throw error

	// You can make new empty slice with make([]TypeData, length, capacity)
	var varStrSliceDecMake = make([]string, 10, 10)

	// Copy another slice to another variable
	var varStrSliceCopy []string
	// copy(destination, source)
	copy(varStrSliceCopy, varStrSlice)

	// Length of array and slice
	fmt.Println(varIntSliceLen, varIntArrayLen)
	fmt.Println(varIntSliceCap, varIntArrayCap)
	fmt.Println(varStrSliceAdd)
	fmt.Println(varStrSliceDecMake)
	fmt.Println(varStrSliceCopy)
}
func main01MapVariableType() {
	// Map is like array or slice
	// In map, you can havae index key with string
	var varStrMap map[string]string // Create new map with key string and value string
	varStrMap = map[string]string{} // You need to initiliaze it first, because the value still nill
	varStrMap["city"] = "Malang"    // You can use string as key and value string
	fmt.Println(varStrMap)

	var varAnyMap map[any]any = map[any]any{} // Create new map with key any variable and value any variable
	varAnyMap[1] = 1
	varAnyMap["first_name"] = "last_name"
	varAnyMap[2.5] = 2.7
	fmt.Println(varAnyMap)

	// Function on map
	var varStrMapDecl map[any]any = make(map[any]any) // You can create map variable like creating slice variable
	fmt.Println(varStrMapDecl)

	// Delete by map key
	delete(varAnyMap, 2.5) // Delete by map key
	fmt.Println(varAnyMap)
}
func main01IfSwitchExpression() {
	var varStr string = "Ach Rozikin"

	/**
	* IF EXPRESSION
	**/
	// If expression will be execute on or more statement if an expression is true
	if true {
		fmt.Println("This is true")
	}

	// Long if expresion
	if false {
		fmt.Println("This is not gonna be print")
	} else if false {
		fmt.Println("This is too")
	} else {
		fmt.Println("This will be printed")
	}

	// Short if statement
	// You can use single expression before if expression
	if text_length := len(varStr); text_length > 5 { // Variable text_length will be only use on this function only
		fmt.Println("Text more than 5")
	} else {
		fmt.Println("Text less than 5")
	}

	/**
	* SWITCH EXPRESSION
	**/
	// Normal switch expression
	switch varStr { // You can use switch expression with string or integer
	case "Ach Rozikin":
		fmt.Println("This is right")
	case "Malang":
		fmt.Println("This is city name")
	case "Jawa Timur":
		fmt.Println("This is province")
	default:
		fmt.Println("Default statement")
	}

	// Short switch statement
	switch text_length := len(varStr); text_length { // text_length variable will be only use in this switch stetement
	case 5:
		fmt.Println("Text lenth is 5")
	case 4:
		fmt.Println("Text lenth is 4")
	case 3:
		fmt.Println("Text lenth is 3")
	default:
		fmt.Println("Text lenth is not 5 4 and 3")
	}

	// Switch without condition : this is like if else but with switch
	switch {
	case varStr == "Ach Rozikin":
		fmt.Println("String is name")
	case varStr == "Malang":
		fmt.Println("String is city")
	case varStr == "Jawa Timur":
		fmt.Println("String is province")
	default:
		fmt.Println("String unidentified")
	}
}

func main01LoopingAndBreakContinueStatement() {
	/**
	* LOOPING STATEMENT
	**/
	// In golang, there is only one looping statement
	// Nomal looping statement
	for i := 0; i < 10; i++ {
		fmt.Println("Print normal loop -", i)
	}

	// For looping with single statement
	counter := 0
	for counter < 10 {
		fmt.Println("Print single statement loop-", counter)
		counter++
	}

	// For range statement to iterate on slice, array, or map
	var varStrMap map[string]string = map[string]string{
		"name":     "Ach Rozikin",
		"city":     "Malang",
		"province": "East Java",
	}
	for index, value := range varStrMap {
		fmt.Println("Index:", index, "; Value:", value)
	}
	/**
	* BREAK CONTINUE
	**/
	// break statement is use for break a looping condition, and stop the looping
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Print normal loop with break -", i)
	}

	// continue statement is use for continue the next looping condition
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("Print normal loop with continue -", i)
	}
}
func main01FunctionStatementWithReturn(name string) string {
	return "Hello World " + name
}
func main01FunctionStatement() {
	// FUNCTION AS VARIABLE
	// You save a function to a variable
	var varFunc func() = func() {
		fmt.Println("This function will print hello world")
	}
	// And then you can call it like a normal function
	varFunc()

	// You can use outside function too
	varFunctOutside := main01FunctionStatementWithReturn("name") // <-- This will be throw error
	fmt.Println(varFunctOutside)

	// But you can't use outside function to a variable if there is no return value
	// varFunctOutsideNoReturn := main01LoopingAndBreakContinueStatement() // <-- This will be throw error

	// FUNCTION AS VARIABLE TO A FUNCTION
	// Or you can pass function to a function with another variable type
	var varFuncNeedFunct func(string, func()) = func(s string, f func()) {
		fmt.Println("Print string input", s)
		f() // Call the given function
	}

	// You pass a string and then pass a function
	varFuncNeedFunct("Ach Rozikin", varFunc)

	// MULTIPLE RETURN VALUE OF FUNCTION
	// A function can return any value
	// This function need string variable, and will return string variable with integer variable
	var varFuncReturnValue func(string) (string, int) = func(s string) (string, int) {
		return s, 1
	}
	// You can get all the value with separated comma between variable
	varStr, vartInt := varFuncReturnValue("a")
	fmt.Println(varStr, vartInt)

	// If you want to ignore the return value, you can use underscore
	varStrNew, _ := varFuncReturnValue("a")
	fmt.Println(varStrNew)

	// NAMED RETURN VALUE
	var varFuncNamedReturnValue func() (string, string) = func() (name, address string) { // You can directly set create name for your return value
		name = "Ach Rozikin"
		address = "Malang"
		return name, address
	}

	varReturnName, varReturnAddress := varFuncNamedReturnValue()
	fmt.Println(varReturnName, varReturnAddress)

	// VARIADIC FUNCTION
	// Variadic function is a function that can receive unlimited ammount of paramater
	// But you just can use it as the last parameter
	var varFuncVariadic func(string, ...int) int = func(s string, i ...int) (total int) {
		fmt.Println("String var :", s)
		for _, v := range i {
			total += v
		}
		return total
	}
	// With variadic function you can use as much as parameter. Or none
	fmt.Println(varFuncVariadic("Hello World", 1, 2, 3, 4))
	fmt.Println(varFuncVariadic("Hello World"))

	// We can use slice of variabel as variadic function parameter too
	var varSliceInt []int = []int{
		1, 2, 3,
	}
	fmt.Println(varFuncVariadic("Hello World", varSliceInt...))

	// ANONYMOUS FUNCTION
	// When we pass a function to a parameter, we can create the function directly with anonymous function
	varFuncNeedFunct("String", func() {
		fmt.Println("Anonymous funct without name")
	})

	// CLOSURE
	// A function can interact with the near or upper variable
	var counter int = 0
	var incrementCount = func() {
		counter++
	}
	incrementCount()
	incrementCount()
	fmt.Println(counter)

	// DEFER PANIC and RECOVER
	// Defer is mechanism of a function after the function is done to be execute
	// Panic is mechanism of function to throw error
	// Recover is mechanism to catch panic
	var varFuncWithDefer func() = func() {
		defer fmt.Println("This will be execute last")
		fmt.Println("This will be execute first")
	}
	varFuncWithDefer()

	// name this function with varFuncWithPanic
	var _ func() = func() {
		defer fmt.Println("varFuncWithPanic : This will be execute last")
		panic("varFuncWithPanic : This will be throw error")
	}
	// varFuncWithPanic() // If you call this, the next code will be not execute

	var varFuncWithRecover func() = func() {
		defer func() {
			fmt.Println("varFuncWithRecover : Please recover")
			panicMessage := recover()
			if panicMessage != nil {
				fmt.Println(panicMessage)
			}
		}() // When nee () on the last, to execute the function
		//  If not, this will be just a pointer
		panic("varFuncWithRecover : This will be throw error")
	}
	varFuncWithRecover()
}

func main01StructDataType() {
	// Struct or I call it structure is a datatype that can contain multiple and nested data type
	// You can directly declare a struct and it's value
	// Create a simple struct
	type address struct {
		city, province string
		postal_code    int
	}

	var varAddress = address{
		city:     "Malang",
		province: "Jawa TImur",
	}
	fmt.Println("varAddress :", varAddress)

	// Or you can declare the value later
	var varAddressDecl address
	varAddressDecl.city = "Surabaya"
	varAddressDecl.province = "Jawa Timur"
	fmt.Println("varAddressDecl :", varAddressDecl)

	// Or you can declare without the field
	var varAddressDecl2 = address{"Jakarta", "DKI Jakarta", 769938}
	fmt.Println("varAddressDecl2 :", varAddressDecl2)

	// Struct field, can be any type or varible. Or it can be struct too
	// Struct person have field address
	type person struct {
		name    string
		address address
	}
	// Declaring multi level struct
	var varRozikin person = person{
		name: "Ach Rozikin",
		address: address{
			city:     "Malang",
			province: "Jawa Timur",
		},
	}
	fmt.Println(varRozikin)
}

/**
* Function Receiver / Method / Function
*
**/
// On golang, you can create function for your data type
type typeInt int

// A function receiver for simple type
func (v typeInt) printTypeInt() {
	fmt.Println("This print called from typeInt receiver. The value : ", v)
}

// Normally, function receiver is use for struct datatype as a method
type typeStruct struct {
}

// A function receiver for simple type
func (t typeStruct) printTypeStruct() {
	fmt.Println("This print called from typeInt receiver. The value : ", t)
}

func main01FunctionReceiver() {
	// To use function receiver/method you need to declare the type first
	// and then you can use the function receiver/method via variable
	// declaring type to a variable
	var varTypeInt typeInt = 1
	var varTypeStruct typeStruct = typeStruct{}

	// Call the function receiver from the variable
	varTypeInt.printTypeInt()
	varTypeStruct.printTypeStruct()
}

/**
* Interface type
*
**/
// Interface is one of the abstract data type
// A BASIC INTERFACE have collection of function/method that not implemented yet
type intfUser interface {
	// string // <-- You can't include type inside a interface.
	// int    // If you include this line, this will be CONSTRAINT INTERFACE
	getName() string
	setName(string)
}

// A interface can contain interface too
type intfUserInside interface {
	intfUser
	addName(string)
}

// STRUCT user implementing intfUser
type user struct {
	name string
	age  int
}

func (u user) getName() string {
	return u.name
}
func (u user) setName(n string) {
	u.name = n
}

// STRUCT animal implementing intfUser
type animal struct {
	name string
}

func (a animal) getName() string {
	return a.name
}
func (a animal) setName(n string) {
	a.name = n
}
func (a animal) addName(n string) {
	a.name = n
}

// func (i intfUser)test(){ <-- You can't use interface to function receiver
//
// }

// Use BASIC INTERFACE a function parameter
func funcImplementationIntfUser(i intfUser) {
	fmt.Println("This called from interface", i.getName())
}

// Interface inside interface as a function parameter
func funcImplementationIntfUserInside(ii intfUserInside) {
	fmt.Println("This called from interface", ii.getName())
}
func main01Interface() {
	// To try the interface, first create variable of implemented interface
	// And it will be user struct and animal struct : user{} animal{}
	var varStructUser = user{
		name: "Rozi",
		age:  26,
	}
	var varStructAnimal = animal{
		name: "Panda",
	}
	// When know that user struct and animal struct is implementing interface intfUser
	// We can pass the struct to function funcImplementationIntfUser
	funcImplementationIntfUser(varStructUser)
	funcImplementationIntfUser(varStructAnimal)
	// Struct animal impement struct under struct too
	funcImplementationIntfUserInside(varStructAnimal)
}

func main01EmptyInterfaceAndNil() {
	/**
	* EMPTY INTERFACE
	**/
	// On golang interface{} mean super parent of all the data type
	// If you declare interface{} as a variable, you can pass anything to that variable
	var varIntfEmpty interface{}
	varIntfEmpty = "Hallo" // Input with string
	fmt.Println(varIntfEmpty)
	varIntfEmpty = 0.1 // Input with float
	fmt.Println(varIntfEmpty)
	varIntfEmpty = 1 // Input with int
	fmt.Println(varIntfEmpty)

	// If you use interface{} as a variable, you can't doing any operation with that variable
	// Before you cast it to another oparateable type
	// fmt.Println(varIntfEmpty + varIntfEmpty) // <-- This will be throw error, because the variable not it
	intValOfVarIntfEmpty, isValidInt := varIntfEmpty.(int)
	if isValidInt {
		fmt.Println("Empty interface valid to cast to int", (intValOfVarIntfEmpty + intValOfVarIntfEmpty))
	} else {
		fmt.Println("Empty interface not a valid int", varIntfEmpty)
	}

	/**
	* NIL VALUE
	**/
	// On some other programming language null is all default value of all variable
	// But, on golang we have nil
	// It's NOT default value for all variable. But it's for some data type like interface, function, map, slice, pointer dan channel
	var varIntfUser intfUser = nil     // Interface data type
	var _ func() = nil                 // Function data type
	var varMap map[string]string = nil // Map data type
	// var varInt int = nil               // <-- You CAN'T use nil on int data type. Int default value is 0
	// varFunc()
	fmt.Println(varIntfUser, varMap)
}

/**
* ERROR INTERFACE
**/
// An error interface is an interface that we can implement to reproduce error for our application
func main01ErrorInterfaceReturnError() (int, error) {
	// If you check errors.new file
	// Is a struct that contain string
	// And implement error interface

	/**

	func New(text string) error {
		return &errorString{text}
	}

	// errorString is a trivial implementation of error.
	type errorString struct {
		s string
	}

	func (e *errorString) Error() string {
		return e.s
	}
	**/

	return 1, errors.New("a")
}
func main01ErrorInterface() {
	// A function return int and error
	// To return error, you need to implement error interface
	varInt, varError := main01ErrorInterfaceReturnError()
	fmt.Println(varInt, varError)
}
func main01TypeAssertions() {
	// Type assertion is function to change the value
	// Of variable interface{}
	var varIntfEmpty interface{}
	varIntfEmpty = "This is string" // Insert string to interface{}
	// We need to cast the interface first to another variable to use it as variable
	varIntOfvarIntfEmpty, varIsValid := varIntfEmpty.(int)
	// If casting valid
	if varIsValid {
		fmt.Println("varIntfEmpty is valid : ", varIntOfvarIntfEmpty)
	} else {
		// If casting not valid
		fmt.Println("varIntfEmpty is not valid : ", varIntfEmpty)
	}

	varIntfEmpty = 1 // Insert int to interface{}
	varIntOfvarIntfEmpty, varIsValid = varIntfEmpty.(int)
	// If casting valid
	if varIsValid {
		fmt.Println("varIntfEmpty is valid : ", varIntOfvarIntfEmpty)
	} else {
		// If casting not valid
		fmt.Println("varIntfEmpty is not valid : ", varIntfEmpty)
	}

	// Try to change the interface value
	// varIntfEmpty = "String"
	// varIntfEmpty = 20

	// We can use type assertion with switch
	switch value := varIntfEmpty.(type) {
	case string:
		fmt.Println("Value is string :", value)
	case int:
		fmt.Println("Value is int :", value)
	default:
		fmt.Println("Value is unknown :", value)
	}
}
func main01PointerDataType() {
	// On programming language, pointer data type is use to
	// pointing the memory address of computer
	// after we got the address, we can change value of the address

	// To demonstrate pointer, first we create a struct
	type person struct {
		name string
	}

	var varStructPerson = person{name: "Rozi"}
	var varStructPerson2 = varStructPerson
	fmt.Println(varStructPerson, varStructPerson2) // Var struct have the same value
	varStructPerson2.name = "Achmad"               // Change the value of varStructPerson2
	fmt.Println(varStructPerson, varStructPerson2) // varStructPerson have different value of varStructPerson

	// Let's try use pointer
	var varStructPersonPointer *person                    // Create a pointer of struct person
	varStructPersonPointer = &varStructPerson             // The pointer take value of memory addres with &
	fmt.Println(varStructPerson, *varStructPersonPointer) // Var struct have the same value. When we use pointer, we need * to get the value of the address
	varStructPersonPointer.name = "Achmad Rozikin"        // Var the value of varStructPersonPointer
	fmt.Println(varStructPerson, *varStructPersonPointer) // varStructPerson and varStructPersonPointer have the same value. Because pointer pointing on varStructPerson

	// Using pointer data type on function
	// First we create a normal function
	var varFuncNormal func(person, string) = func(p person, s string) {
		p.name = s
	}
	// Create a pointer function
	var varFuncPointer func(*person, string) = func(p *person, s string) {
		p.name = s
	}

	// Let's try to change the value of struct person
	fmt.Println("Current value : ", varStructPerson)
	varFuncNormal(varStructPerson, "Edited with varFuncNormal") // This is will not give any effect
	fmt.Println("After change the value : ", varStructPerson)

	// With function pointer
	fmt.Println("Current value : ", varStructPerson)
	varFuncPointer(&varStructPerson, "Edited with varFuncPointer") // This will change the value
	fmt.Println("After change the value : ", varStructPerson)

	// We can create a null pointer value with keyword new
	var varStructPersonPointerNew = new(person) // Equivalent to var varStructPersonPointer *person
	var varStructPersonPointerNew2 = varStructPersonPointerNew
	varStructPersonPointerNew2.name = "Name is change"
	fmt.Println(*varStructPersonPointerNew, *varStructPersonPointerNew2) // The 2 value will have the same value

	// We can use pointer on function method too
	var varStructWithMethodPointer typeStructWithMethodPointer = typeStructWithMethodPointer{name: "Achmad Rozikin"}
	fmt.Println("Current value : ", varStructWithMethodPointer)
	varStructWithMethodPointer.setName("Achmad Rozikin Edited With setName")
	fmt.Println("After change the value : ", varStructWithMethodPointer)

	// Change value with pointer
	fmt.Println("Current value : ", varStructWithMethodPointer)
	varStructWithMethodPointer.setNameWithPointer("Achmad Rozikin Edited With setNameWithPointer")
	fmt.Println("After change the value : ", varStructWithMethodPointer)
}

// A struct
type typeStructWithMethodPointer struct {
	name string
}

// Function receiver without pointer
func (t typeStructWithMethodPointer) setName(s string) {
	t.name = s
}

// Function receiver with pointer
func (t *typeStructWithMethodPointer) setNameWithPointer(s string) {
	t.name = s
}
