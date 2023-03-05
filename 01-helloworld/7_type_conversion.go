package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	a := 21
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 42
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Go 에서는 cast라고 안하고 conversion이라고 함
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var f float64 = 30.0
	var ff float64 = float64(a)

	fmt.Println(f)
	fmt.Println(f + ff + .5)
	fmt.Println(int(f + ff + .5))
}
