package main

import ("fmt")

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{324, 456, 768, 87798}
	x = append(x, y...)
	fmt.Println(x)
}