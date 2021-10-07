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

	/*
		Comparison Operator

		==  equal operator, checks if values on both side of the op. are equal
		!=  not equal operator, checks if values on both sides are NOT Equal
		>   greater than, checks if left operand is greater than the right one
		<   lesser than, checks if left operand is lesser than the right one
		>=  greater than or equal to, checks if left operand is greater or equal to than the right one
		<=  lesser than or equal to, checks if left operand is lesser or equal to than the right one

	*/

	p, q, r := 5, 5, 7

	fmt.Printf("%v == %v is %v\n", p, r, p == r)
	fmt.Printf("%v != %v is %v\n", p, r, p != r)
	fmt.Printf("%v > %v is %v\n", p, q, p > q)
	fmt.Printf("%v < %v is %v\n", p, r, p < r)
	fmt.Printf("%v >= %v is %v\n", p, q, p >= q)
	fmt.Printf("%v <= %v is %v\n", p, r, p <= r)

	/*
		Logical Operators

		&&  logical and , eg. true && true => true, true && false => false
		||  logical or , eg. true || true => true, true || false => true
		!   logical negation , eg. !true => false, !false => true

	*/

	i, j, k := 2, 4, 5

	// 2 == 4 is false , 2 < 5 is also true, false and(&&) true is false
	fmt.Printf("(%v == %v)&&(%v < %v) is %v\n", i, j, i, k, (i == j) && (i < k))

	// 4 >= 4 is true , 5 <= 4 is flase, false or(||) true is true
	fmt.Printf("(%v >= %v)||(%v <= %v) is %v\n", j, j, k, j, (j >= j) || (k <= j))

	// 2 == 2 is true, not(!) true is false
	fmt.Printf("!(%v == %v) is %v \n", i, i, !(i == i))

	// 2 > 5 is false, not(!) flase is true
	fmt.Printf("!(%v > %v) is %v \n", i, k, !(i > k))

}

/*

OUTPUT:
-------
Addition of 8 and 5 is 13
Difference between 6.6 and 4.7 is 1.8999999999999995
Multiplication of 8 and 5 is 40
Division of 6.6 and 4.7 is 1.404255319148936
Remainder when 8 is divided by 5 is 3
8 is of type int, while 6.6 is of type float64
Adding 8 and 6.6 by converting 6.6 into int is 14
Dividing 5 and 6.6 by converting 5 into float is 1.3199999999999998
Value of x is 5
Incremented value of x is 6
Decremented value of x is 5
Value of y is 15
y+= 17
y-= 12
y/= 2
y*= 6
y%= 1
5 == 7 is false
5 != 7 is true
5 > 5 is false
5 < 7 is true
5 >= 5 is true
5 <= 7 is true
(2 == 4)&&(2 < 5) is false
(4 >= 4)||(5 <= 4) is true
!(2 == 2) is false
!(2 > 5) is true

*/
