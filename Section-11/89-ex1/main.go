package main

import ("fmt")

func main() {
	x := [5]int{5, 7, 9, 56, 34}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}