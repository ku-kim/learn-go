package main

import "fmt"

func main() {
	f := returnFunc()

	fmt.Println(f())
	fmt.Println(returnFunc()()) // 신기하군.!
	fmt.Printf("%T", f)

}

func returnFunc() func() int {
	return func() int {
		return 42
	}
}
