package main

import "fmt"

type person struct {
	first string
	last string
	favFlavors []string
}

func main()  {
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

	m := map[string]person{
		p1.last:p1,
		p2.last:p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("_-------_")
	}
}