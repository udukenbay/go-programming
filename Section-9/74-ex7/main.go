package main

import ("fmt")

func main() {
	x := "asd"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDONBONDBONDOND", x)
	} else {
		fmt.Println("neither")
	}
}