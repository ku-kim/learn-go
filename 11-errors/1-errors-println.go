package main

import "fmt"

/*
err를 항상 처리하자!
*/
func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n) // 6 (h e l l o \n)
}
