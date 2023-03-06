package main

import "fmt"

func main() {
	foo()
	bar("haha")
	s1 := woo("Money")
	fmt.Println(s1)

	x, y := twoReturn("a", "b")
	fmt.Println(x, y)
}

func twoReturn(s string, s2 string) (string, bool) {
	a := fmt.Sprint(s, s2)
	b := false
	return a, b
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func woo(s string) string {
	return fmt.Sprint("Hello from woo ", s)
}

// Go에서는 매개변수가 기본으로 pass by value
func bar(s string) {
	fmt.Println("hello", s)
}

func foo() {
	fmt.Println("hello from foo")
}
