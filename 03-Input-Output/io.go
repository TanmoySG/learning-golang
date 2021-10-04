package main

// The fmt Package is Used Primarily for I/O Operations in Golang
import "fmt"

func main() {

	midWord := "World"
	name, age := "Golang", 11

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
	*/
	fmt.Printf("")
}
