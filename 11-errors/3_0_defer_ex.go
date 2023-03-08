package main

import "fmt"

func main() {
	var x int // 0
	x++
	fmt.Println("x: ", x)

	y := c()
	fmt.Println("y: ", y)

	fmt.Println("defer는 후입 선출이다")
	deferTest()
}

func deferTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 1,2,3,4가 아닌 4,3,2,1
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1 // 1 -> defer 함수 i++ -> 2
}
