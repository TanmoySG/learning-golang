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

## Loops - <kbd>for</kbd>

Golang has only one Looping Statement - <kbd>for</kbd> 

The <kbd>for</kbd> loop has the following syntax -

```
for <initiation>; <test-condition>; <post statement> {
	<do task>
}
```

- initation involves declaring a temporary variable to start the program
- test-condition involves testing of a condition based on which the loop will run
- post statement runs after the task inside the braces are completed.
	
#### Program to Print 0 to 4

```
for i := 0; i < 5; i++ { // i is a temporary variable
	fmt.Printf("%v ", i)
}
```

#### Program to print upto 6 fibonacci numbers

```
fib := []int{0, 1}

for i := 2; i <= 5; i++ {
	fib = append(fib, fib[i-1]+fib[i-2])
}

fmt.Printf("Fibonacci Series upto 6 numbers: %v \n", fib)
```

There is no while loop in go. But a for loop can also be used as a while loop in Go using the following Syntax.

```
i := 1 // initiation

for <condition> {
	<do something>
	<post statement>  // increament
}
```
#### Program to print upto 6 fibonacci numbers using a while-like for loop

```
fib = []int{0, 1}

i := 2 // initiation

for i <= 5 { // test condition
	fib = append(fib, fib[i-1]+fib[i-2])
	i++ // post statement
}

fmt.Printf("Fibonacci Series upto 6 numners: %v \n", fib)
```

#### Program to count backwards from 10 to 0 with 2 jumps

Required Output: <kbd>10 -> 8 -> 6 -> 4 -> 2 -> 0</hbd>
```
for num := 10; num >= 0; num = num - 2 {
	fmt.Printf("%v ", num)
}
```

#### for loop with multiple conditions

```
for i, j := 10, 20; j > 0; i, j = i-1, j-5 {
	fmt.Printf("%v-%v ", i, j)
}
```

### range in for-loop

for-loop <kbd>range</kbd> iterates over elements in a variety of data structures. The range returns the key/index and the value at that key/index.


#### <kbd>range</kbd> for Arrays

```
nums := []int{2, 3, 4}
for index, num := range nums { // iterates over nums
	fmt.Printf("%v at %v \n", num, index)
}
```

#### <kbd>range</kbd> for Strings

```
str := "Golang"
for index, char := range str { // iterates over str
	fmt.Printf("%v at %v \n", char, index) // prints the ASCII equivalent
}
```

#### <kbd>range</kbd> for Map

```
dict := map[string]int{
	"First":  1,
	"Second": 2,
	"Third":  3,
}

for key, value := range dict { // iterates over dict
	fmt.Printf("%v at %v \n", value, key)
}

fmt.Printf("Keys : \n")

for key := range dict { // iterates over dict, returns only indices
	
    fmt.Printf("%v ", key)
}

fmt.Printf("\nValues : \n")

for _, value := range dict {
	fmt.Printf("%v ", value)
}
```

## Continue Statement

Continue skips the remaining code in the loop, and directly goes to the post-statement (increament) and to the begining of the loop.
	
```
fmt.Println("continue Demo -")

// Program to skip the number 3 while printing others

for i := 1; i < 6; i = i + 1 {
	if i == 3 {
		fmt.Println("Encountered 3")
		continue // skips code after this line and goes to the start of the loop
	}
	fmt.Printf("%v, %v*2=%v, %v+2=%v\n", i, i, i*2, i, i+2)
}
```

## Break Statement

Break terminates its parent loop/funtion, i.e. Like continue, break too skips the code, but unlike continue, it doesn't go to the begining of the loop, but comes out of it.

```
fmt.Println("break Demo -")

// Program to terminate operations when number 3 is encountered

nums = []int{1, 2, 3, 4, 5, 6, 7}

for _ , i := range nums {
	if i == 3 {
		fmt.Println("Encountered 3")
		break // breaks out of loop
	}
	fmt.Printf("%v, %v*2=%v, %v+2=%v\n", i, i, i*2, i, i+2)
}
```