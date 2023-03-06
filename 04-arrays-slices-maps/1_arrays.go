package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)

	x[2] = 42
	fmt.Println(x)

	fmt.Println(len(x))

	fmt.Println(x[:3])

	y := x // value copy
	fmt.Println(y)
	y[0] = 44
	fmt.Println(y, x) // 서로 다름
}
