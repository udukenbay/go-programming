package main

import "fmt"

//enclosing a variable so that we limit its scope

func main() {
	var x int
	fmt.Println(x)
	x++

	{
		y := 42
		fmt.Println(y)
	}

	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("hello")
}
