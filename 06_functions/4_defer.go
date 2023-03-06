package main

import "fmt"

func main() {
	deferTest()
}

func deferTest() {
	fmt.Println("start")
	defer fmt.Println("defer1") // 제일 마지막 실행
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	fmt.Println("end")
}
