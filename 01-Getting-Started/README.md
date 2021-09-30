# Getting Started with Go

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software, and is developed by **Robert Griesemer**, **Rob Pike** & **Ken Thompson** at **Google**. Find more about Go at [golang.org](https://golang.org/).

## Structure of a Go Program

A Go Executable file is defined by using the <kbd>package main</kbd> at the start of the go file.

A Go program can be structured as follows
```
package main

import "fmt"

func main(){
	fmt.Println("Hello World!");
}
```

## Go Commands

Go Commands are used to compile, run and build go programs.

The <kbd>go</kbd> keyword is used to execute commands on the terminal.

### go run 

<kbd>go run</kbd> - Used to Compile and run a Go Application.
```
go run main.go
```

### go build 

<kbd>go build</kbd> - Used to Compile and generate a Go Application executable file.
```
go build main.go
```

<kbd>go build -o name</kbd> - Used to Compile and generate a Go Application executable file with custom name.
```
go build -o main.go app
```

### Compiling Go Application for Different Platforms

Go Programs can be compiled and bundled for cross-platform usage using the <kbd>GOOS</kbd> and <kbd>GOARCH</kbd> specifications.

<kbd>GOOS</kbd> is used to specify the platform for which the Go App need to be Compiled. This can have the values - *windows*, *linux*, *darwin*, etc.

<kbd>GOARCH</kbd> is used to specify the OS Architecture for the specified target OS. This can have values - *amd64*, *arm64*, *wasm*, etc.

```
GOOS=windows GOARCH=amd64 go build -0 winapp.exe
GOOS=linux GOARCH=amd64 go build -0 linuxapp
GOOS=darwin GOARCH=amd64 go build -0 macapp
```

Read more about GOOS and GOARCH [here](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63).