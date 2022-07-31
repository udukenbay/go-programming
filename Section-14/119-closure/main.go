package main

import "fmt"

//enclosing a variable so that we limit its scope

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
}

func foo() {
	fmt.Println("hello")
	x++
}
