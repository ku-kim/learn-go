package main

import "fmt"

//var x = 2

const (
	_ = iota
	//kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	basic()

	//kb := 1024
	//mb := kb * 1024
	//gb := mb * 1024

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

func basic() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	var z int = y << 1
	fmt.Printf("%d\t\t%b\n", z, z)
}
