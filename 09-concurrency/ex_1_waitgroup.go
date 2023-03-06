package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func main() {

	wg2.Add(2)

	go func() {
		fmt.Println("Hey")
		wg2.Done()
	}()

	go func() {
		fmt.Println("Hi")
		wg2.Done()
	}()

	wg2.Wait()
}
