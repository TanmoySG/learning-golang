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

*/

func main() {
	var i string // variable i of type string
	i = "World"  // assigning value to i

	var j int = 26 // variable j of type int initalized with value 26

	fmt.Println(i)
	fmt.Println(j)
}
