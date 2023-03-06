package main

import "fmt"

func main() {
	//x := type{value} // composite literal
	x := range_test()
	copy_test(x)

	append_test()
	append_delete_test()

	make_test()
}

func make_test() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	y := make([]int, 1, 1)
	fmt.Println(len(y), cap(y)) // 1 1

	y = append(y, 1)
	fmt.Println(len(y), cap(y)) // 2 2

	y = append(y, 2, 3)
	fmt.Println(len(y), cap(y)) // 4 4

	y = append(y, 4)
	fmt.Println(len(y), cap(y)) // 5 8  // capacity가 2배씩 증가

}

func append_delete_test() {
	x := []int{4, 5, 6, 7, 8}
	y := []int{99, 999, 999}
	x = append(x, y...) // ...으로 모든 값 append
	fmt.Println(x, len(x), cap(x))

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func append_test() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)
	x = append(x, 111)
	fmt.Println(x)

	y := []int{99, 999, 999}
	x = append(x, y...) // ...으로 모든 값 append
	fmt.Println(x)
}

func copy_test(x []int) {
	y := x[:] // 포인터 위치만 넘기기 때문에 동일
	y[2] = 42
	fmt.Println(x, y)

	yy := x // 포인터 위치만 넘기기 때문에 동일, slice는 모두 포인터 위치, 여기서 array와 다름
	yy[0] = 42
	fmt.Println(x, yy)
}

func range_test() []int {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
	return x
}
