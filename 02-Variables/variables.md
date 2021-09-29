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
