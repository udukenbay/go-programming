package main

import ("fmt")

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny", "bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
		fallthrough
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")		
	}
}