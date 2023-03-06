package main

import "fmt"

// & : 해당 값의 주소
// * : 주소의 값
func main() {

	a := 42
	fmt.Println(a)
	fmt.Println(&a)        // & 주소값 확인
	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int, pointer type

	//var b int = &a // error
	var b *int = &a // *int(포인터) 타입으로 지정해야함
	c := &a         // error
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(*b)
	fmt.Println(*&a) // 이렇게는 잘 안쓰겠지만..

	// 값 변경
	*b = 43
	fmt.Println(a)
	fmt.Println(*b)
}
