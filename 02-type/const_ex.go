package main

import "fmt"

const a = 42
const b = "hello world"

const (
	aa         = 1
	bb float64 = 11.1
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	//a = 11 // compile error
	fmt.Println(aa)
	fmt.Println(bb)

}
