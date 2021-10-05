package main

// The fmt Package is Used Primarily for I/O Operations in Golang
import "fmt"

func main() {

	midWord := "World"
	name, age := "Golang", 11
	height := 5.56
	arr := []int{1, 2, 3, 4}
	boolVar := true
	pi := 3.145968

	// fmt.Println()

	// The fmt.Println() function writes to the Standard Output.
	fmt.Println("Hello World!")

	/*
		Println() adds a New Line at the end of each line.
		Println() takes multiple arguments and concatenates
		each argument while adding white-space between them.
	*/
	fmt.Println("Hello", midWord, "I am", name, "and I am", age, "years old.")

	// fmt.Printf()

	// The fmt.Printf() function is used to print formatted outputs to the standard output.
	fmt.Printf("%s is %d years old.\n", name, age)

	/*
		fmt.Printf() uses specifications called "Verbs" to format the output.
		These verb can be used to specify the value of the variable, the type of the variable
		as well as the format in which the variable value will be printed.
		It doesn't add newline at the end of the output.

		These verbs are used as placeholders in predefined output strings that are replaced
		with the values. Verbs are used with the % (percentage-sign) followed by the required
		format and type of verb.
	*/

	fmt.Printf("Value %s of type %T\n", name, name)

	/*
		Common Verbs used for Formatting are:
		%d - decimal
		%f - float
		%s - string
		%q - double quoted strings
		%v - any value/type
		%T - Type of the Value/variable
		%t - boolean
		%p - pointer
		%c - char/rune represented by the corresponding Unicode Value
		%b - Base 2 (Binary Value)
		%x - hexadecimal
	*/

	// %f placeholder-verb for float64 values and %T for Type
	fmt.Printf("This is a %T variable with value %f\n", height, height)

	// %d for int and %T for Type
	fmt.Printf("This is an %T variable with value %d\n", age, age)

	// %b is used for displaying the Base2-Binary Equivalent of the assigned variable
	fmt.Printf("This is an %T variable with value %b in base2\n", age, age)

	// %x is used for displaying the Hexadecimal Equivalent of the assigned variable
	fmt.Printf("This is an %T variable with value %x in Hex-Base16\n", age, age)

	// %s for string type 
	fmt.Printf("This is an %T variable with value %s\n", name, name)

	// %v is used for holding any values irrespective of its type 
	// and hence can hold any value without specifying the type verb
	fmt.Printf("This is an %T variable with value %v\n", arr, arr)
	fmt.Printf("This is an %T variable with value %v\n", midWord, midWord)

	// While %t is used to hold Boolean type Values
	// %T is used to hold the type of the value
	fmt.Printf("This is an %T variable with value %t\n", boolVar, boolVar)

	/*
		For Floating point Variables whose values have trailing digits after decimal point
		e.g. 2.36579 - 0.36579 are the trailing digits, there might be a requirement for only
		displaying the number upto a certain decimal places. Eg. 2.36579 can be displayed upto
		2 decimal places to make 2.36. 
	
		This can be achieved by using the %0.2f verb, where the 0.2 specifies that the number
		is to be displayed till 2 decimal places. For n-decimal places we use %0.nf.
	*/
	fmt.Printf("%f upto 2 decimal places = %0.2f\n", pi, pi)
	fmt.Printf("%f upto 3 decimal places = %0.3f\n", pi, pi)
	fmt.Printf("%f upto 4 decimal places = %0.4f\n", pi, pi)

}
