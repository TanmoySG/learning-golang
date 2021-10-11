# Control Flow 

## Conditional Statements

Conditional Statements are used to decide the action to be taken based on a condition.

Conditional Statements in Golang are:

- If-Else
- Swith-Case

### If Else Conditional Statement

If-Else block can have multiple conditional to decide the action. The first condition is handled by the If statement while, the next conditions are handled by the Else-If block and if none of the conditions are satisfied then a default else block is executed.

```
package main

import "fmt"

func main() {
	a, b := 25, 59

	if a < b { // condition 1
		fmt.Println("Condition 1 is True.")
	} else if a > b { // condition 2
		fmt.Println("Condition 2 is True.")
	} else { // when both conditions fail
		fmt.Println("Both Conditions were False.")
	}
}
```
While we can use a variable with a boolean value in languages like python for a condition in if statements, we cannot do so in Golang.

```
c := true
if c {
	//do something
}
```
The above code raises error.


### Switch-Case

Switch-Case in Go works same as that of its counterpart in C/C++. A switch block is initialized with the condition that has to be matched with a case inside the Switch block. The <kbd>case</kbd> keyword is used to specify the different cases and actions if the case matches, it also has a default case that executed when no case matches.

```
language := "javascript"

switch language {
case "python":
	fmt.Println("Language is Python")
case "go":
	fmt.Println("Go Rocks!")
default:
	fmt.Println("Some other awesome language")
}

OUTPUT:
Some other awesome language
```

## Loops

Golang has only one Looping Statement - <kbd>for</kbd>