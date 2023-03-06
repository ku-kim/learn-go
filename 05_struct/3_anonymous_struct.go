package main

import "fmt"

func main() {
	p3 := struct {
		first, last string
		age         int
	}{first: "ku",
		last: "kim",
		age:  30}

	fmt.Println(p3)
}
