package main

import "fmt"

// 함수 표현식은 일급 객체, 타입으로 처리됨, Function types
func main() {
	f := func() {
		fmt.Println("my first func expression")
	}

	f()
}
