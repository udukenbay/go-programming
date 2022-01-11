package main

import ("fmt")

func main() {
	xs := []string{"James", "Bond", "Shaken, not stirred"}
	xt := []string{"Miss", "Moneypenny", "Hellooooo, James."}
	x := [][]string{xs, xt}
	fmt.Println(x)

	for i, xa := range x {
		fmt.Println("record:", i)
		for j, val := range xa {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, val)
		}
	}
}