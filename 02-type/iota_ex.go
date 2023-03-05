package main

import "fmt"

const (
	x = iota
	y = iota
	z
)

const (
	xx = iota
	yy = iota
	zz
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// 초기화 됨
	fmt.Println(xx)
	fmt.Println(yy)
	fmt.Println(zz)
}
