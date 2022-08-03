package main

import "fmt"

//interface is allow us to define behavior.
//& also allow us to do polymorphism.

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (s person) speak() {
	fmt.Println("I am", s.first, s.last, " - the person speak")
}

//keyword identifier type
type human interface {
	speak()
}

//interface says: "Hey babe, if you got this method(speak in this ex.) or these methods, then you are my TYPE"

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar person", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar secretAgent", h.(secretAgent).first)
	default:
		fmt.Println("I was passed into bar", h)
	}
}

//So any other type that has the method speak is also of type human
//A value can be of more than one type******
//sa1 has two types: secretAgent and human(as it has speak method)

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"MoneyPenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	//convetsion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x) //here we have conversion
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
