package main

import (
	"fmt"
)

/*

Variables in Golang

-----------------------------------
- var variable_name type          -
- variable_name = value           -
--------------- OR ----------------
-                                 -
- var variable_name type = value  -
-                                 -
--------------- OR ----------------
-                                 -
- variable_name := value          -
-                                 -
-----------------------------------

var - a keyword to declare a variable
variable_name - name of the variable
type - type of the declared variable
value - value assigned to the variable

The  :=  is a short variable declaration operator that
doesn't require initialization with type of the variable
and by default assigns a type based on the value. But this
operation cannot be used outside a funtion.

Variables can also be declared on the package level,
i.e. outside funtions.  

Mvltiple Variables can also be declared together as a variable block.
-----------------------------------
-  var (                          -
-     variable1 type = value      -
-     variable2 type              -
-     variable3 type              -
-    )                            -
-----------------------------------

*/

// Declaring variable outside funtion/ on package level
var varOut string = "This variable was declared outside function." 

// Declaring multiple variables together using single 'var' block
var (
	variable1 int = 21
	variable2 string
	variable3 float32
)

func main() {
	var i string // variable i of type string
	i = "World"  // assigning value to i

	var j int = 26 // variable j of type int initalized with value 26

	k := 78.5 // k initialized with value 78.5 and assigning type float64 by default based on assigned value

	variable2 = "Var Block Entity"
	variable3 = 3.33

	fmt.Println(varOut)
	fmt.Printf("%v of type %T \n", i, i)
	fmt.Printf("%v of type %T \n", j, j)
	fmt.Printf("%v of type %T \n", k, k)

	fmt.Printf("variable1 = %v was declared and initalized inside var block \n", variable1)
	fmt.Printf("variable2 = %v was declared inside var block and value assigned in main function \n", variable2)
	fmt.Printf("variable3 = %v was declared inside var block and value assigned in main function \n", variable3)
}
