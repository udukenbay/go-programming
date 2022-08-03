package main

import "fmt"

func main() {
	fmt.Println("Hello")

	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	f1 := func(x int) {
		fmt.Println("the year stars watching", x)
	}
	f1(42)
}
