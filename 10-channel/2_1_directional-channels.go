package main

import "fmt"

func main() {
	// 송신 전용 채널(receive)
	cr := make(chan<- int, 2)
	fmt.Printf("%T\n", cr)
	cr <- 43
	// fmt.Println(<-c) // 수신 불가
	// fmt.Println(<-c)

	// 수신 전용 채널(send)
	cs := make(<-chan int)
	fmt.Printf("%T\n", cs)

	// 송수신 전용채널은 타입 변경 불가
	// (chan int)(cs) // error
}
