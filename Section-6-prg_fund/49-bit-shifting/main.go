package main

import ("fmt")

const (
	_ = iota
	kb = 1 << (iota * 10)//1024
	mb = kb * 1024
	gb = 1024 * mb
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)


	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}