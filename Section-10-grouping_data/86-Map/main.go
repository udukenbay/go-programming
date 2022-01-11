package main

import ("fmt")

func main() {
	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barbaras"])
	
	v, ok := m["Barbaras"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE PRINT", v)	
	}

	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}
}
