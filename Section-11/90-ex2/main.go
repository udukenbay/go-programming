package main

import ("fmt")

func main() {
	x := []int{3, 5, 6, 7, 2, 1, 2, 3, 4, 5}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}