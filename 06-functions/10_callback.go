package main

import "fmt"

// callback은 인수로 func를 전달하는 것
// 함수는 일급객체, 함수 반환, 변수에 함수 할당, 인자로 함수 넘기기 가능
func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println(s2)

	s3 := odd(sum, ii...)
	fmt.Println(s3)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}
