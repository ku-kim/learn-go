package main

import "fmt"

var xx int
var yy string
var zz bool

var xxx int = 42
var yyy string = "James Bond"
var zzz = true

type T1 int

func main() {
	fmt.Println("\nex1")
	//ex1()

	fmt.Println("\nex2")
	//ex2()

	fmt.Println("\nex3")
	//ex3()

	fmt.Println("\nex4")
	ex5()

}

func ex5() {
	var x T1 = 10
	x = 20

	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func ex3() (int, error) {
	result := fmt.Sprintf("%v\t%v\t%v", xxx, yyy, zzz)
	return fmt.Println(result)
}

func ex2() {
	fmt.Println(xx)
	fmt.Println(yy)
	fmt.Println(zz)
}

func ex1() {
	var x = 42
	y := "James Bond"
	var z bool = true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
