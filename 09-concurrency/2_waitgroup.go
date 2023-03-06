package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo2()
	bar2()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

}

func foo2() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar2() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar: ", i)
	}
}
