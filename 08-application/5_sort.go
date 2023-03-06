package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 2, 5, 522, 124, 3, 3, 66612}
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("=======\n")
	xs := []string{"a", "bbbz", "1"}
	sort.Strings(xs)
	fmt.Println(xs)
}
