package main

import "fmt"

//defer
func main()  {
	defer foo()
	bar()
	beer()
	mid()
}

func foo()  {
	fmt.Println("i have defer")
}

func bar() {
	fmt.Println("no defer")
}

func beer()  {
	defer func() {
		fmt.Println("beers defer ran")
	}()
	fmt.Println("beer")
}

func mid()  {
	fmt.Println("i am in the middle")
}