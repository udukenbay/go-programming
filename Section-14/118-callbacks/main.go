package main

import "fmt"

//A callback is passing a func as an argument

//so func  we could: return functions-117
//we could: assign functions to variables-116
//we could: pass a function into another function as an argument-118
func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("All numbers:", s)

	s2 := even(sum, ii...)
	fmt.Println("Even numbers:", s2)

	s3 := odd(sum, ii...)
	fmt.Println("Odd numbers:", s3)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

//func(xi ...int) int callback function as an argument which passed into even function
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
