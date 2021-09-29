package main

import (
	"fmt"
)

/*

1. Variables in Golang
-----------------------------------
- var variable_name type          -
- variable_name = value           -
--------------- OR ----------------
-                                 -
- var variable_name type = value  -
-                                 -
-----------------------------------

var - a keyword to declare a variable
variable_name - name of the variable
type - type of the declared variable
value - value assigned to the variable


2. The  :=  is a short variable declaration operator that
doesn't require initialization with type of the variable
and by default assigns a type based on the value. But this
operation cannot be used outside a funtion.
-----------------------------------
-                                 -
- variable_name := value          -
-                                 -
-----------------------------------


3. Variables can also be declared on the package level,
i.e. outside funtions.
-----------------------------------
-                                 -
- var variable_name type = value  -
-                                 -
- func main() {                   -
-	// Function Body              -
- }                               -
-                                 -
-----------------------------------


4. Multiple Variables can also be declared together as a variable block.
-----------------------------------
-  var (                          -
-     variable1 type = value      -
-     variable2 type              -
-     variable3 type              -
-    )                            -
-----------------------------------


5. Re-declaration of variables
Variables can be redeclared in different scopes. The "Innermost Scope"
takes precedence over other Scopes. This is called "Shadowing". For Example:
--------------------------------------------
- var scopedVar int = 25 // Scope-1        -
- func main(){                             -
-	// Prints '25' as variable is          -
-	// declared outside funtion            -
-	// scope and is the innermost          -
-	// scope till this line is Scope-1     -
-	fmt.Printf("\nPrinted inside Scope-1: %v", scopedVar)                 -
-                                          -
-	// Re-declaration of Variable          -
-	// Making this the innermost scope     -
-	// and overriding the Scope-1          -
-                                          -
-	var scopedVar int = 58 // Scope-2      -
-                                          -
-	// Prints '58' as variable is          -
-	// declared in Scope-2, which is       -
-	// is the innermost scope till.        -
-	fmt.Printf("\nPrinted inside Scope-2: %v", scopedVar)                 -
- }                                        -
--------------------------------------------
- OUTPUT:                                  -
- 25  // Printed for Scope-1               -
- 58  // Printed for Scope-2               -
--------------------------------------------
Scope-1 is "Package Level Scope"
Scope-2 is "Function Level Scope"


6. All declared variables in Golang HAVE TO BE USED.
----------------------------------
- var i int = 25                 -
- j:= 56 // unused               -
- fmt.Println(i)                 -
----------------------------------
- Error: j declared and not used -
----------------------------------


7. Type conversion between numeric types.
Variables type can be converted in numeric types
using explicit type conversion.
---------------------------------------------------------
- var i int = 25                                        -
- fmt.Println("%v of type %T", i, i)                    -
- var j float32                                         -
- j = i // invalid - implicit type conversion           -
- j = float32(i)  // valid - explicit type conv         -
- fmt.Println("%v of type %T after conversion", i, i)   -
---------------------------------------------------------
- OUTPUT                                                -
- 25 of type int                                        -
- 25 of type float32 after conversion                   -
---------------------------------------------------------
While type converting, data might be lost. For Example
converting float 83.3 into int results in loss of data
after decimal i.e. 0.3 is lost.
Eg: 83.3 (float32) -(int)-> 83 (int)


8. Type conversion between numeric and string types is not
avaiable through general type conversion. Eg:
-----------------------------------
- var i int = 42                  -
- fmt.Println(i)                  -
- var j string = string(i)        -
- fmt.Println(j)                  -
-----------------------------------
- OUTPUT:                         -
- 42                              -
- *                               -
-----------------------------------
"*" is printed because converting int to string by general
method converts the number into its UNICODE equivalent.

To solve this, Go has a special built-in string conversion package
- strconv [ https://pkg.go.dev/strconv@go1.17 ]

To convert int to string use strconv.Itoa() method that converts
int to its ASCII String Equivalent
--------------------------------------
- import ( "fmt" , "strconv")        -
- func main(){                       -
-   var i int = 42                   -
-   var j string = strconv.Itoa(i)   -
-   fmt.Println(j)                   -
- }                                  -
--------------------------------------
- OUTPUT:                            -
- 42                                 -
- 42                                 -
--------------------------------------


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
