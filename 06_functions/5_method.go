package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// method, secretAgent 구조체의 메서드
func (s secretAgent) speak() {
	fmt.Println("ltk : ", s.ltk)
}

func main() {
	sa1 := secretAgent{
		person: person{first: "ku",
			last: "kim"},
		ltk: true,
	}

	fmt.Println(sa1)

	sa1.speak()
}
