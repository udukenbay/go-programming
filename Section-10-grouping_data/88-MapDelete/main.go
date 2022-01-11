package main

import ("fmt")

func main() {
	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	
	fmt.Println(m)
}
