# Golang Data Types

Data types specify the type of data that a valid Go variable can hold. 

## Integer Type

In Go, both signed and unsigned integers are available in four different sizes - 8, 16, 32, 64 Bits. The signed int is represented by <kbd>int</kbd> and the unsigned integer is represented by <kbd>uint</kbd>.

```
// Unsigned Integer - uint, uint8, uint16, uint32, uint64
var unsignedInteger uint = 12 

fmt.Printf("This is an %T with value %v\n", unsignedInteger, unsignedInteger)

// Signed Integer - int, int8, int16, int32, int64
var signedInteger int = 15   

fmt.Printf("This is an %T with value %v\n", signedInteger, signedInteger)
```

## Floating Point Data

Floating point numbers are represented using <kbd>float32</kbd> and <kbd>float64</kbd> for 32 bit and 64 bit floating point values, respectively.

```
// 32 bit floating point
var floatingPoint32Bit float32 = 3.1459 

fmt.Printf("This is an %T with value %v\n", floatingPoint32Bit, floatingPoint32Bit)

// 64 bit floating point
var floatingPoint64Bit float64 = 1.1699 

fmt.Printf("This is an %T with value %v\n", floatingPoint64Bit, floatingPoint64Bit)
```

## String Type Variables

Collection of Characters are called Strings. It is represented using <kbd>string</kbd> keyword.

```
var stringVariable string = "Quick Brown Fox" // A string

fmt.Printf("This is an %T with value %q\n", stringVariable, stringVariable)
```

## Boolean Variables

Boolean is represented by <kbd>bool</kbd>. 
```
var boolVariableTrue bool = true   // boolean with true value

var boolVariableFalse bool = false // boolean with false value

fmt.Printf("This is an %T with value %t\n", boolVariableTrue, boolVariableTrue)

fmt.Printf("This is an %T with value %t\n", boolVariableFalse, boolVariableFalse)
```

## Arrays

Arrays are collection of similar type of data having a fixed size. It is represented using - the array size between Square Brackets followed by the data-type followed by elements between curly braces - <kbd>[array_size]data_type{array_element, array_element}</kbd>

```
var arrayVariable = [4]int{1, 5, 15, 100} // array of fixed size 4

fmt.Printf("This is an %T with value %v\n", arrayVariable, arrayVariable)
```

## Slice

A slice is a collection of similar type of data having a variable size/length. It is represented using the same syntax as that of an array, but without the array size - <kbd>[]data_type{array_element, array_element}</kbd>

```
var sliceVariable = []int{11, 52, 415, 560} // slice

fmt.Printf("This is an %T with value %v\n", sliceVariable, sliceVariable)
```

## Maps

A Map is key-value paired data with unique key - simillar to Python Dictionaries. It is defined using the <kbd>map</kbd> keyword -
```
map[data_type_of_key ]data_type_of_value{
	key1: value1,
	key2: value2,
}
```
This can be used as -
```
var mapVariable = map[string]int{
	"Python": 30,
	"Go":     11,
} // map of key type string and value type int

fmt.Printf("This is an %T with value %v\n", mapVariable, mapVariable)
```

## Struct Data Type

It is a user-defined data type with sequence of named elements. It is defined using the <kbd>struct</kbd> keyword.

```
type structureVar struct{ a, b int }
structVar := structureVar{1, 2} // Initialize a object of type structureVar with values 1, 2

fmt.Printf("This is an %T with value %v\n", structVar, structVar)
```

## Pointers

Pointer is a variable that stores the memory address of another variable. 

<kbd>&</kbd> - **Address of** Operator. Eg. <kbd>ptr := &a</kbd> is same as <kbd>ptr = address of a</kbd>

<kbd>*</kbd> - **Value at Address** Operator. Eg. <kbd>val := *ptr</kbd> is same as <kbd>val = value at address in ptr</kbd>

```
var sourceVar int = 2    // variable who's address is required to be stored

// & is the "address of" operator so, pointerVar stores the address of sourceVar
pointerVar := &sourceVar 

fmt.Printf("This is an %T with value %v\n", pointerVar, pointerVar) // Address of the source variable

// the * operator is used to display the value at the address stored in the pointer variable
fmt.Printf("This is an %T with value %v\n", *pointerVar, *pointerVar) // Value stored at the address stored in pointerVar
```