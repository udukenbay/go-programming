package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main()  {
	t := truck {
		fourWheel: true,
		vehicle: vehicle {
			doors: 4,
			color: "white",
		},
	}

	s := sedan {
		luxury: true,
		vehicle: vehicle {
			doors: 2,
			color: "blue",
		},
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}