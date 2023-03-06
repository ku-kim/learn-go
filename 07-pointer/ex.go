package main

import "fmt"

type person struct {
	first, last string
}

func (p *person) changeMe(first, last string) {
	p.first = first
	p.last = last

}

func changeMe2(p *person) {
	p.last = "aa"
	p.first = "kk"
}

func main() {
	ex1()

	p := person{first: "kukim", last: "kim"}

	fmt.Println(p.first, p.last)
	p.changeMe("ha", "hi")
	fmt.Println(p.first, p.last)

	fmt.Println(p.first, p.last)
	changeMe2(&p)
	fmt.Println(p.first, p.last)

}
func ex1() {
	i := 42
	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println("======\n")
}
