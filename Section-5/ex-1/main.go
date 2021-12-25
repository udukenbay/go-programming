package main

import ("fmt")

//short declaration x := 1

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
}