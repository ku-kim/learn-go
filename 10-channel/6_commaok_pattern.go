package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(eve, odd, quit)

	receive(eve, odd, quit)

}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
	// close(e)
	// close(o)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from eve channel: ", v)
		case v := <-o:
			fmt.Println("from odd channel: ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok : ", i, ok) // i: value(false), ok : 채널 닫혀있는지 여부(닫혀있으면 false, 열려있으면 true)
				return
			} else {
				fmt.Println("from comma ok: ", i)
			}
		}
	}
}
