package main

import "fmt"

/*
채널에서 송수신이 이뤄져야함.
채널을 차단해야함 (channel block!)

현재 수신자가 없어서 멈춰버림
*/
func main() {
	c := make(chan int)

	go func() { // 동시성에서 한쪽이 채널로 데이터 쏨
		c <- 42
	}()

	fmt.Println(<-c)
}
