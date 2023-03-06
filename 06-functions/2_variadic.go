package main

import "fmt"

func main() {
	result := sum_test(1, 2, 3, 4, 5)

	fmt.Println("sum : ", result)

}

func sum_test(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // slice로 들어옴

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
