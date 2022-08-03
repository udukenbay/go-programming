package main

import "fmt"

func main()  {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := foo(ii...)
	fmt.Println("This is variadic parameter", n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7}
	x := bar(ii2)
	fmt.Println("This is total", x)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}