package main

import "fmt"

func main() {
	a, b, c, d := 8, 5, 6.6, 4.7

	/*
		Arithmetic Operators are used to perform math-operations and calculations

		Go supports the following Arithmetic Operators -

		+  Sum
		-  Difference
		*  Multiplication
		/  Divide
		%  Remainder/Modulus

		Go doesn't have any power operator so, math.Pow() is used.

	*/

	// sun
	fmt.Printf("Addition of %v and %v is %v\n", a, b, a+b)

	// difference
	fmt.Printf("Difference between %v and %v is %v\n", c, d, c-d)

	// multiplication
	fmt.Printf("Multiplication of %v and %v is %v\n", a, b, a*b)

	// division
	fmt.Printf("Division of %v and %v is %v\n", c, d, c/d)

	// remainder
	fmt.Printf("Remainder when %v is divided by %v is %v\n", a, b, a%b)

	/*
		Go is a strongly typed language, hence it doesn't support
		arithmatic operations between values of different types.

		a, b := 5, 5.6
		fmt.Println(a+b)

		The above raises an error as a is int while b is float.

		Hence, type-conversion is required for such cases.

	*/

	// arithmatic ops between difference datatypes using type-conversion
	fmt.Printf("%v is of type %T, while %v is of type %T\n", a, a, c, c)
	fmt.Printf("Adding %v and %v by converting %v into int is %v\n", a, c, c, a+int(c))
	// While converting Types, we see some data is lost, 6.6 becomes 6
	fmt.Printf("Dividing %v and %v by converting %v into float is %v\n", b, c, b, c/float64(b))
	// Here data is added while converting, 8 becomes 8.0

	/*
		Increment and Decrement Operators

		x := 1

		++ increment operator increases the value by 1, eg. x++ => 2
		-- decrement operator decreases the value by 1, eg. x-- => 0
	*/

	x := 5
	fmt.Printf("Value of x is %v\n", x)
	x++
	fmt.Printf("Incremented value of x is %v\n", x) // increment
	x--
	fmt.Printf("Decremented value of x is %v\n", x) // decrement

	/*
		Assignment Operator

		y := 6

		=   simple assignement
		+=  increment and assign, eg. y+=2 => y = y+2 => 8
		-=  decrement and assign, eg. y-=2 => y = y-2 => 4
		*=  multiply and assign, eg. y*=2 => y = y*2 => 12
		/=  divide and assign, eg. y/=2 => y = y/2 => 3
		%=  modulus and assign	, eg. y%=2 => y = y%2 => 0
	*/

	var y int

	y = 15
	fmt.Printf("Value of y is %v\n", y)
	y += 2
	fmt.Printf("y+= %v\n", y) // increment and assign
	y -= 5
	fmt.Printf("y-= %v\n", y) // decrement and assign
	y /= 6
	fmt.Printf("y/= %v\n", y) // divide and assign
	y *= 3
	fmt.Printf("y*= %v\n", y) // multiply and assign
	y %= 5
	fmt.Printf("y%%= %v\n", y) // modulus and assign


}
