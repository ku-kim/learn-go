package main

import "fmt"

func main() {

	if true {
		fmt.Println("hello")
	}

	if x := 42; x == 42 {
		fmt.Println(x)
	}

	x := 40
	if x == 42 {
		fmt.Println(x)
	} else if x == 41 {
		fmt.Println(x)
	} else {
		fmt.Println(x)
	}
}
