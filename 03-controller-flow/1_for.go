package main

import "fmt"

func main() {

	basic_for()
	nested_for()
	while_for()
	inf_for()
	string_range()

}

func string_range() {
	for pos, char := range "collection" {
		fmt.Println(pos, char)
	}
}

func inf_for() {
	x := 1
	for {
		if x == 3 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")
}

func while_for() {
	x := 1
	for x < 3 {
		fmt.Println(x)
		x++
	}
}

func nested_for() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}

func basic_for() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
