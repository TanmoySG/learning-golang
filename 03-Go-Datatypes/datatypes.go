package main

import "fmt"

// Data Types in Go

func main() {

	// integer - Signed and Unsigned
	var unsignedInteger uint = 12 // uint, uint8, uint16, uint32, uint64
	var signedInteger int = 15    // int, int8, int16, int32, int64

	fmt.Printf("This is an %T with value %v\n", unsignedInteger, unsignedInteger)
	fmt.Printf("This is an %T with value %v\n", signedInteger, signedInteger)

	// float - Floating Point Data
	var floatingPoint32Bit float32 = 3.1459 // 32 bit floating point
	var floatingPoint64Bit float64 = 1.1699 // 32 bit floating point

	fmt.Printf("This is an %T with value %v\n", floatingPoint32Bit, floatingPoint32Bit)
	fmt.Printf("This is an %T with value %v\n", floatingPoint64Bit, floatingPoint64Bit)

	// strings - Group of Characters
	var stringVariable string = "Quick Brown Fox" // A string

	fmt.Printf("This is an %T with value %q\n", stringVariable, stringVariable)

	// boolean - True or False value
	var boolVariableTrue bool = true   // boolean with true value
	var boolVariableFalse bool = false // boolean with false value

	fmt.Printf("This is an %T with value %t\n", boolVariableTrue, boolVariableTrue)
	fmt.Printf("This is an %T with value %t\n", boolVariableFalse, boolVariableFalse)

	// arrays - collection of similar type of data having a fixed size
	var arrayVariable = [4]int{1, 5, 15, 100} // array of fixed size 4

	fmt.Printf("This is an %T with value %v\n", arrayVariable, arrayVariable)

	// slice - collection of similar type of data having a variable size/length
	var sliceVariable = []int{11, 52, 415, 560} // slice

	fmt.Printf("This is an %T with value %v\n", sliceVariable, sliceVariable)

	// maps - a key-value paired data with unique key - simillar to Python Dictionaries
	var mapVariable = map[string]int{
		"Python": 30,
		"Go":     11,
	} // map of key type string and value type int

	fmt.Printf("This is an %T with value %v\n", mapVariable, mapVariable)

	// struct - is a user-defined data type with sequence of named elements
	type structureVar struct{ a, b int }
	structVar := structureVar{1, 2} // Initialize a object of type structureVar with values 1, 2

	fmt.Printf("This is an %T with value %v\n", structVar, structVar)

	// pointer - Pointer is a variable that stores the memory address of another variable
	var sourceVar int = 2    // variable who's address is required to be stored

	// & is the "address of" operator so, pointerVar stores the address of sourceVar
	pointerVar := &sourceVar 
	
	fmt.Printf("This is an %T with value %v\n", pointerVar, pointerVar) // Address of the source variable

	// the * operator is used to display the value at the address stored in the pointer variable
	fmt.Printf("This is an %T with value %v\n", *pointerVar, *pointerVar) // Value at the address stored in pointerVar


}
