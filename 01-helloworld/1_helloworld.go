package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")

	foo()

	//for i := 0; i < 100; i++ {
	//	if i%2 == 0 {
	//		fmt.Println("짝수 %d", i)
	//	}
	//}

	bar()

	n, err := fmt.Println("12345")

	fmt.Println(n, err)
}

func bar() {
	fmt.Println("bar")
}

func foo() {
	fmt.Println("Foo")
}
