package main

import "fmt"

/*
포인터 언제쓰는가?
- 값 변경할 때

why?

Go는 모든 것을 Pass by Value로 전달함 (사본으로) (중요!)
-> reference 직접 수정 불가


*/

func main() {
	step1()

	step2()
}

func step2() {
	x := 0
	foo2(&x)
	fmt.Println(x) // 값 변경됨
	fmt.Println("=====\n")
}

func foo2(y *int) {
	fmt.Println(y)  // 주소
	fmt.Println(*y) // 0
	*y = 42
	fmt.Println(y)  // 주소
	fmt.Println(*y) // 42
}

func step1() {
	x := 0
	foo(x)         // pass by value
	fmt.Println(x) // 값변경되지 않음

	fmt.Println("=====\n")
}

func foo(y int) {
	fmt.Println(y)
	y = 42
	fmt.Println(y)
}
