package main

import "fmt"

// Closure : 뭔가 닫는다는 개념
// 변수 범위 코드의 한 영역으로 제한되도록 하는 것
// 변수 범위 제한

var x int // 전체 범위(전역변수, scope : global)

func main() {
	fmt.Println(x) // 0
	x++
	fmt.Println(x) // 1
	foo()
	fmt.Println(x) // 2

	{
		y := 42
		fmt.Println(y)
	}
	//fmt.Println(y) // error, scope : code block

	// closure로 범위 제한할 수 있음
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(a()) // 4
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2
}

func incrementor() func() int {
	var xx int
	return func() int {
		xx++
		return xx
	}
}

func foo() {
	fmt.Println("foo")
	x++
}
