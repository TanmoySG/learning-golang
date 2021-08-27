package main

import (
	"fmt"
)

/*

Re-declaration of variables

Variables can be redeclared in different scopes.
The "Innermost Scope" takes precedence over other Scopes. 

This is called "Shadowing". 

*/

var scopedVar int = 25 // Scope-1 : Package Scope

func main() {
	// Prints '25' as variable is declared outside funtion
	// scope and is the innermost scope till this line is Scope-1
	fmt.Printf("%v - Package Scope\n", scopedVar)

	// Re-declaration of Variable Making this the innermost scope
	// and overriding the Scope-1
	var scopedVar int = 58 // Scope-2 : Funtion Scope

	// Prints '58' as variable is declared in Scope-2, 
	// which is  is the innermost scope till.
	fmt.Printf("%v - Function Scope (innermost scope) \n", scopedVar)
}
