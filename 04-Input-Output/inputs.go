package main

import (
	"fmt"
	"os"
)

// User Inputs from Command Line

/*
	To get command line inputs from the Command line, the user
	needs to pass the argument(CL Arguments) while running the go file

	go run file_name.go argument1 argument2 argumentn

	CL Inputs are used using the OS Package in go standard library.
*/

func main() {

	/*
		var firstInput int
		var secondInput string
	*/
	/*
		os.Args

		The os.Args method returns an array of arguments.
		The first argument in the array is the path of the Go file.
		Subsequent elements in the array are the passed arguments
	*/

	fmt.Println("os.Args returned: ", os.Args)

	fmt.Println("os.Args First argument is the Path: ", os.Args[0])
	fmt.Println("os.Args Other Arguments: ", os.Args[1:])
	fmt.Println("Number of arguments user entered:  ", len(os.Args))


}
