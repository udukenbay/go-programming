package main

import ("fmt")

func main() {
	switch {
	case false:
		fmt.Println("should not pring")
	case true:
		fmt.Println("should print")
	}
}