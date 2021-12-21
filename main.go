package main

import ("fmt")

var y = 43
//DECLARE there is a VARIABLE with the IDENTIFIER "z"
//& that the VARIABLE with the IDENTIFIER "z" is of TYPE int
//ASSIGNS the ZERO VALUE of of TYPE int to "z"
var z int

func main() {
	//short declaration operator
	//DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
