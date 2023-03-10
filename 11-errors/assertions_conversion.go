package main

import (
	"fmt"
)

type hotdog int

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	// assertions 예제
	foo2(customErr{
		info: "need more coffee",
	})

	// conversion 예제
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x) // conversion
	fmt.Println(y)
	fmt.Printf("%T", y)
}

func foo2(e error) {
	fmt.Println("foo ran -", e, "\n", e.(customErr).info) // assertions
}
