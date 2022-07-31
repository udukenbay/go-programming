package main

import "fmt"

//A callback is passing a func as an argument

//so func  we could: return functions-117
//we could: assign functions to variables-116
//we could: pass a function into another function as an argument-118
func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}
