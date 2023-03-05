package main

import "fmt"

var x = 42
var z = "hello"
var a = `aaa "hello" bb`

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // %T : type 확인

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
