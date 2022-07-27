package main

import ("fmt")

// create your own type
// call it "UNDERLYING TYPE"

type newType int
var x newType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}