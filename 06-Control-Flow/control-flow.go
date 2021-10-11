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

	/*
		While we can use a variable with a boolean value in languages like
		python for a condition in if statements, we cannot do so in go

		c := true
		if c {
			//do something
		}

		This raises error.
	*/

	language := "javascript"

	/*
		Switch-Cases
	*/

	switch language {
	case "python":
		fmt.Println("Language is Python")
	case "go":
		fmt.Println("Go Rocks!")
	default:
		fmt.Println("Some other awesome language")
	}

	// complex switch conditions
	switch n := 5; true {
	case n > 4:
		fmt.Println("Greater than 4")
	case n == 5:
		fmt.Println("Equal to 5")
	case n < 4:
		fmt.Println("Smaller than 4")
	}

	/* 

	Go To Statements

	*/
}
