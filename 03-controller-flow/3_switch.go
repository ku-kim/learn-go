package main

import "fmt"

func main() {
	switch_notexist_expr()
	switch_expr()
}

func switch_expr() {
	n := 2
	switch n {
	case 2:
		fmt.Println("n", n)
	}
}

func switch_notexist_expr() {
	switch {
	case false:
		fmt.Println("출력 X")
	case 1 == 1:
		fmt.Println("1=1")
		fallthrough // switch 에 기본적으로 break가 포함되어있음, 따라서 아래 case까지 모두 포함하려면 fallthrough 사용
	case 2 == 2:
		fmt.Println("2==2")
	default:
		fmt.Println("default")
	}
}
