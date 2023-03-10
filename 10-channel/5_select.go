package main

import "fmt"

// https://velog.io/@moonyoung/golang-channel-with-select-%ED%97%B7%EA%B0%88%EB%A6%AC%EB%8A%94-%EC%BC%80%EC%9D%B4%EC%8A%A4-%EC%A0%95%EB%A6%AC%ED%95%98%EA%B8%B0
func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	receive(eve, odd, quit)

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from eve channel: ", v)
		case v := <-o:
			fmt.Println("from odd channel: ", v)
		case v := <-q:
			fmt.Println("from quit channel: ", v)
			return
		}
	}
}
