# Variables in GoLang

A variable is name given to a storage area that the programs can manipulate.

### Data Types in Go

The common types of data a Go Variable can have are:

|              Type              |                           Description                           |
| :----------------------------: | :-------------------------------------------------------------: |
| uint8, uint16, uint32, uint164 |               Unsigned 8, 16, 32, 64-bit integers               |
|   int8, int16, int32, int64    |                Signed 8, 16, 32, 64-bit integers                |
|        float32, float64        |             32-bit & 64-bit floating-point numbers              |
|     complex64, complex128      | Complex numbers with float32 & float64 real and imaginary parts |

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

