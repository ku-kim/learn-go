package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	for v := range c { // close() 되기 전까지 채널에서 range 반복함
		fmt.Println(v)
	}

	fmt.Println("exit ddd")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // rance 붙이지 않으면 rance 빠져나올 수 없음
}
