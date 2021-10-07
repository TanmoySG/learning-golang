# Operators in Go

Operators allow us to perform different kinds of operations on operands. Operators are classified into,

- Arithmetic Operator
- Increment and decreament Operator
- Assignment Operator
- Comparison Operator
- Logical Operator

## Arithmetic Operator

Arithmetic Operators are used to perform math-operations and calculations

Go supports the following Arithmetic Operators -

<kbd>+</kbd>  Sum

<kbd>-</kbd>  Difference

<kbd>*</kbd>  Multiplication

<kbd>/</kbd>  Divide

<kbd>%</kbd>  Remainder/Modulus

Go doesn't have any power operator so, math.Pow() is used.

```
// sum
fmt.Printf("Addition of %v and %v is %v\n", a, b, a+b)

// difference
fmt.Printf("Difference between %v and %v is %v\n", c, d, c-d)

// multiplication
fmt.Printf("Multiplication of %v and %v is %v\n", a, b, a*b)

// division
fmt.Printf("Division of %v and %v is %v\n", c, d, c/d)

// remainder
fmt.Printf("Remainder when %v is divided by %v is %v\n", a, b, a%b)
```

Go is a strongly typed language, hence it doesn't support arithmatic operations between values of different types.

```
a, b := 5, 5.6
fmt.Println(a+b)
```

The above raises an error as a is int while b is float. Hence, type-conversion is required for such cases.

```
// arithmatic ops between difference datatypes using type-conversion
fmt.Printf("%v is of type %T, while %v is of type %T\n", a, a, c, c)
fmt.Printf("Adding %v and %v by converting %v into int is %v\n", a, c, c, a+int(c))
// While converting Types, we see some data is lost, 6.6 becomes 6
fmt.Printf("Dividing %v and %v by converting %v into float is %v\n", b, c, b, c/float64(b))
// Here data is added while converting, 8 becomes 8.0
```

## Increment and decreament Operator

Increaments or Decrements value by 1. 

<kbd>++</kbd> increment operator increases the value by 1, eg. <kbd>x := 1; x++ => 2</kbd>

<kbd>--</kbd> decrement operator decreases the value by 1, eg. <kbd>x := 1; x-- => 0</kbd>

```
x := 5
fmt.Printf("Value of x is %v\n", x)
x++
fmt.Printf("Incremented value of x is %v\n", x) // increment
x--
fmt.Printf("Decremented value of x is %v\n", x) // decrement
```

## Assignment Operator

Used to assign value to a variable. 

<kbd>=</kbd>   simple assignement

<kbd>+=</kbd>  increment and assign

<kbd>-=</kbd>  decrement and assign

<kbd>*=</kbd>  multiply and assign

<kbd>/=</kbd>  divide and assign

<kbd>%=</kbd>  modulus and assign

```
var y int

y = 15 // assignment operator
fmt.Printf("Value of y is %v\n", y) // 15
	
y += 2 // increment and assign
fmt.Printf("y+= %v\n", y)  // 17

y -= 5 // decrement and assign
fmt.Printf("y-= %v\n", y) // 12

y /= 6 // divide and assign
fmt.Printf("y/= %v\n", y) // 2

y *= 3 // multiply and assign
fmt.Printf("y*= %v\n", y) // 6

y %= 5 // modulus and assign
fmt.Printf("y%%= %v\n", y) // 1
```

## Comparison Operator

<kbd>==</kbd>  equal operator, checks if values on both side of the op. are equal

<kbd>!=</kbd>  not equal operator, checks if values on both sides are NOT Equal

<kbd>></kbd>  greater than, checks if left operand is greater than the right one

<kbd><</kbd>   lesser than, checks if left operand is lesser than the right one

<kbd>>=</kbd>  greater than or equal to, checks if left operand is greater or equal to than the right one

<kbd><=</kbd>  lesser than or equal to, checks if left operand is lesser or equal to than the right one

```
p, q, r := 5, 5, 7

fmt.Printf("%v == %v is %v\n", p, r, p == r)  // false
fmt.Printf("%v != %v is %v\n", p, r, p != r)  // true
fmt.Printf("%v > %v is %v\n", p, q, p > q)    // false
fmt.Printf("%v < %v is %v\n", p, r, p < r)    // true
fmt.Printf("%v >= %v is %v\n", p, q, p >= q)  // true
fmt.Printf("%v <= %v is %v\n", p, r, p <= r)  // true
```

## Logical Operators

<kbd>&&</kbd>  logical and , eg. <kbd>true && true => true, true && false => false</kbd>

<kbd>||</kbd>  logical or , eg. <kbd>true || true => true, true || false => true</kbd>

<kbd>!</kbd>   logical negation , <kbd>eg. !true => false, !false => true</kbd>

```
i, j, k := 2, 4, 5

// 2 == 4 is false , 2 < 5 is also true, false and(&&) true is false
fmt.Printf("(%v == %v)&&(%v < %v) is %v\n", i, j, i, k, (i == j) && (i < k))     // false

// 4 >= 4 is true , 5 <= 4 is flase, false or(||) true is true
fmt.Printf("(%v >= %v)||(%v <= %v) is %v\n", j, j, k, j, (j >= j) || (k <= j))   // true

// 2 == 2 is true, not(!) true is false
fmt.Printf("!(%v == %v) is %v \n", i, i, !(i == i)) // false

// 2 > 5 is false, not(!) flase is true
fmt.Printf("!(%v > %v) is %v \n", i, k, !(i > k))   // true
```