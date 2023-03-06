package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println(p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{name: "kukim"}
	p.speak()

	saySomething(&p)
}
