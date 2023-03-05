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

	//b = a // error

}
