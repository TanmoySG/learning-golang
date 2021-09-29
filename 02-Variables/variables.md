# Variables in GoLang

A variable is name given to a storage area that the programs can manipulate.

### Data Types in Go

The common types of data a Go Variable can have are:

|             Type              |                                       Description                                        |
| :---------------------------: | :--------------------------------------------------------------------------------------: |
| uint8, uint16, uint32, uint64 |                           Unsigned 8, 16, 32, 64-bit integers                            |
|   int8, int16, int32, int64   |                            Signed 8, 16, 32, 64-bit integers                             |
|       float32, float64        |                          32-bit & 64-bit floating-point numbers                          |
|     complex64, complex128     |             Complex numbers with float32 & float64 real and imaginary parts              |
|             byte              | A single byte represents ASCII characters. Golang does not have any data type of ‘char’. |
|            string             |                      String is a read only slice of bytes in golang                      |
|             bool              |             The data type is bool and has two possible values true or false              |

## Declaring Variables in Go

Variables are declared in Go using the <kbd>var</kbd> keyword followed by the **variable-name** and <kbd>data-type</kbd>.

```
var variable_name type   
variable_name = value 
```

Variables can also be declared by initializing it with a value.

```
var variable_name type = value
```

### Declaring variables using <kbd>:=</kbd> Operator

Variables can be declared using the <kbd>:=</kbd> - **Walrus Operator**. 

```
variable_name := value
```

The  <kbd>:=</kbd> operator is a short variable declaration operator that doesn't require initialization with type of the variable and by default assigns the type based on the value. But this operation cannot be used outside a funtion.

### Package Level Variable declaration

Variables can also be declared on the package level,
i.e. outside funtions.

```
var variable1 int32 = 5  

func main() {                   
    // Function Body
}
```

### Declaring Multiple Variables Together

Multiple Variables can also be declared together as a **variable block**.

```
var (
    variable1 int32 = 56
    variable2 float64
    variable3 bool = false
)
```

### Re-declaration of Variables

Variables can be redeclared in different scopes. The "Innermost Scope" takes precedence over other Scopes. This is called "Shadowing". For Example:

```
var scopedVar int = 25 // Scope-1 

func main(){                      
	// Prints '25' as variable is   
	// declared outside funtion     
	// scope and is the innermost   
	// scope till this line is Scope-1
	fmt.Printf("\nPrinted inside Scope-1: %v", scopedVar)          
                                   
	// Re-declaration of Variable   
	// Making this the innermost scope
	// and overriding the Scope-1 

	var scopedVar int = 58 // Scope-2 
                                   
	// Prints '58' as variable is   
	// declared in Scope-2, which is
	// is the innermost scope till. 
	fmt.Printf("\nPrinted inside Scope-2: %v", scopedVar)          
} 

OUTPUT:
Printed inside Scope-1: 25
Printed inside Scope-2: 58
```
In the above program,
- Scope-1 is "Package Level Scope"
- Scope-2 is "Function Level Scope"

### All declared variables in Golang ***Have to be used***

Golang Checks if a declared variable is used or not while compiling. If a variable is defined/declared but not used it throws an error.

```
var i int = 25
j:= 56 // unused
fmt.Println(i)

OUTPUT:
Error: j declared but not used
```

### Numeric Type COnversion

Variable type can be changed in numeric data-type variables using explicit type conversion.

```
var i int = 25
fmt.Println("%v of type %T", i, i)
var j float32
j = i // invalid - implicit type conversion
j = float32(i)  // valid - explicit type conv
fmt.Println("%v of type %T after conversion", i, i)

OUTPUT:
25 of type int
25 of type float32 after conversion   
```
While type converting, data might be lost. For Example converting float 83.3 into int results in loss of data after decimal i.e. 0.3 is lost.

Eg: 83.3 (float32) --(int)-> 83 (int)

### Type Conversion between numeric and not numeric data-types

Type conversion between numeric and string types is not
avaiable through general type conversion. For Example:

```
var i int = 42
fmt.Println(i)
var j string = string(i)
fmt.Println(j)

OUTPUT:
42
*
```

"*" is printed because converting int to string by general
method converts the number into its UNICODE equivalent.

To solve this, Go has a special built-in string conversion package - [strconv](https://pkg.go.dev/strconv@go1.17)

To convert int to string use <kbd>strconv.Itoa()</kbd> method that converts int to its ASCII String Equivalent.

```
import ( "fmt" , "strconv")
func main(){ 
    var i int = 42  
    var j string = strconv.Itoa(i) 
    fmt.Println(j) 
} 

OUTPUT: 
42 
42 
```

## A Sample Program for variables

```
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
```