package main

import "fmt"

/*

There are many ways to print to console/terminal in go.
This is provided by the "fmt" package. The 3 main ways are:

1. fmt.Print()
2. fmt.Printf()
3. fmt.Println()

1. fmt.Print()
Print formats using the default formats for its operands
and writes to standard output. Spaces are added between
operands when neither is a string. It returns the number
of bytes written and any write error encountered.
---------------------------------------------------
-  fmt.Print(name, " is ", age, " years old.\n")  -
---------------------------------------------------

2. fmt.Printf()
Printf formats according to a format specifier and writes
to standard output. It returns the number of bytes written
and any write error encountered.
-----------------------------------------------------
-  	fmt.Printf("%s is %d years old.\n", name, age)  -
-----------------------------------------------------

3. fmt.Println()
Println formats using the default formats for its operands
and writes to standard output. Spaces are always added between
operands and a newline is appended. It returns the number of
bytes written and any write error encountered.
-------------------------------------------------
-  	fmt.Println(name, "is", age, "years old.")  -
-------------------------------------------------

String Formatting Verbs
Some of the general-purpose string formatting verbs that are
used most often are:

- %v    the value in a default format
	    when printing structs, the plus flag (%+v) adds field names
- %#v   a Go-syntax representation of the value
- %T    a Go-syntax representation of the type of the value
- %%    a literal percent sign; consumes no value

Strings can be formatted using fmt.Printf() as
--------------------------------------------------------------------------
- fmt.Printf("%v is the value and %T is the type.", // Format of string  -
-             variable_name, variable_name)  // Variables                -
--------------------------------------------------------------------------
In the above code, the value is printed in-place of %v, and the Type of
variable is printed in-place of %T.

*/

func main() {

	var name string = "Go"
	var age int = 11

	fmt.Print(name, " is ", age, " years old.\n")  // fmt.Print()
	fmt.Printf("%s is %d years old.\n", name, age)  // fmt.Printf()
	fmt.Println(name, "is", age, "years old.")  // fmt.Println()


	// Formatted String with String Formatting Verbs
	var i int = 25
	fmt.Printf("Variable i has value %v and is of type %T\n", i, i)
}
