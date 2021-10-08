package main

import "fmt"

func main() {

	a, b, c := 25, 59, true

	_ = c
	/*
		If-Else Conditional Statements
	*/

	if a < b { // condition 1
		fmt.Println("Condition 1 is True.")
	} else if a > b { // condition 2
		fmt.Println("Condition 2 is True.")
	} else { // when both conditions fail
		fmt.Println("Both Conditions were False.")
	}

}
