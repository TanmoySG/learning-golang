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
- var variable_name type = value  -
-----------------------------------

var - a keyword to declare a variable
variable_name - name of the variable
type - type of the declared variable
value - value assigned to the variable while declaration

*/

func main() {
	var i string // variable i of type string
	i = "World"  // assigning value to i

	var j int = 26 // variable j of type int initalized with value 26

	fmt.Println(i)
	fmt.Println(j)
}
