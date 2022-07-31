package main

import "fmt"

//enclosing a variable so that we limit its scope

func main() {
	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("hello")
}
