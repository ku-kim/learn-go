package main

import "fmt"

// 익명함수, 함수 이름 없음
func main() {

	func() {
		fmt.Println("Anonymous Func ran")
	}()

	func(x int) {
		fmt.Println("Anonymous Func ran", x)
	}(42)

	// func expression
	f := func(str string) {
		fmt.Println("Anonymous Func ran", str)
	}
	f("haha")
}
