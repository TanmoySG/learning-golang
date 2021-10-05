# Golang Data Types

Data types specify the type of data that a valid Go variable can hold. 

## Integer Type

In Go, both signed and unsigned integers are available in four different sizes - 8, 16, 32, 64 Bits. The signed int is represented by <kbd>int</kbd> and the unsigned integer is represented by <kbd>uint</kbd>.

```
// Unsigned Integer - uint, uint8, uint16, uint32, uint64
var unsignedInteger uint = 12 

// Signed Integer - int, int8, int16, int32, int64
var signedInteger int = 15   

fmt.Printf("This is an %T with value %v\n", unsignedInteger, unsignedInteger)

fmt.Printf("This is an %T with value %v\n", signedInteger, signedInteger)
```
