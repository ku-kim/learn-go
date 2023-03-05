package main

import "fmt"

var y int = 42

func main() {
	fmt.Printf("%T\n", y)  // 타입
	fmt.Printf("%b\n", y)  // 2진수
	fmt.Printf("%x\n", y)  // 16진수
	fmt.Printf("%#x\n", y) //0x가 있는 16진수
}
