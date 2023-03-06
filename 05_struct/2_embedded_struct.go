package main

import "fmt"

type person2 struct {
	first string
	last  string
}

type secretAgent struct {
	person2 // embedded struct
	ltk     bool
}

func main() {
	sa1 := secretAgent{person2{"James", "Bond"}, true}

	fmt.Println(sa1)
	fmt.Println(sa1.person2.first, sa1.last, sa1.ltk)
	// sa1.person2.first 또는 sa.last 처럼 접근할 수 있음, 승격된다고 함
}
