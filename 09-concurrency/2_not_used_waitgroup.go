package main

import (
	"fmt"
	"runtime"
)

// main 함수 실행되기전에 초기화는 함수도 있음
func init() {

}

/*
동시성 패키지 : runtime
https://pkg.go.dev/runtime

아래 코드는 foo 출력되지 않음
why?

main 실행되는 고루틴 하나 있는데, foo()의 고루틴 (동시성) 실행되기 전에 이미 프로세스가 종료되었기 때문에 보이지 않는 것임
-> 이를 위해서 동기화, WaitGroup이 필요함

	sync 패키지
*/
func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	go foo() // 왜 foo는 출력되지 않을까?
	bar()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
	}
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar: ", i)
	}
}
