package main

import "fmt"

//In Golang, the defer keyword is used to delay the execution of a function or a statement until the nearby function returns.
//In simple words, defer will move the execution of the statement to the very end inside a function.

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
