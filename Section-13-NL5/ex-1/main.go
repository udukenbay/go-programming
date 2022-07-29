package main

import ("fmt")

type person struct {
	first string
	last string
	favFlavors []string
}

func main() {
	p1 := person {
		first: "Indira",
		last: "Kabdoldina",
		favFlavors: []string{
			"chocolate",
			"martini",
			"melon",
		},
	}

	p2 := person {
		first: "Urziya",
		last: "Dukenbay",
		favFlavors: []string{
			"strawberry",
			"kiwi",
			"watermelon",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}
	
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}