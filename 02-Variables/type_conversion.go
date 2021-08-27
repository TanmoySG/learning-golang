package main

import (
	"fmt"
	"strconv"
)

/*

1. Type conversion between numeric types.
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


2. Type conversion between numeric and string types is not
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

func main() {

	// Numeric to Numeric Type Conversion
	fmt.Println("Numeric Conversion")
	var i int = 42
	fmt.Printf("%v, %T \n", i, i)
	var j string = string(i)
	fmt.Printf("%v, %T \n", j, j)

	// Numeric to String Type Conversion
	fmt.Println("String Conversion")
	var k int = 42
	fmt.Printf("%v, %T \n", k, k)
	var l string = strconv.Itoa(k)
	fmt.Printf("%v, %T \n", l, l)
}
