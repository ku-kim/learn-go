package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex // mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // lock
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // unlock
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("counter: ", counter)
}
