package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	result := sum_test2(x...) // 넣을 떄 ... 으로 넣어야함 호오..

	fmt.Println("sum : ", result)

	a := "prefix"
	strings := []string{"a", "b", "c"}
	s := slice_string(a, strings...)
	fmt.Println(s)
}

func slice_string(a string, strings ...string) string {
	fmt.Println(a)
	fmt.Println(strings)
	return fmt.Sprint(a, strings)
}

func sum_test2(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // slice로 들어옴

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
