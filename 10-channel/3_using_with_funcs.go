package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c) // 이곳에 go 붙이면, 끝나지 않기 때문에(waitGroup 없음) 정상 작동하지 않음

	fmt.Println()
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
