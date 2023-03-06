package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var counter = 0

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)

	// 4,5는 mutex, atomic으로 위 문제 해결하기 (풀진 않음)
}
