package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
}

func main() {

	ex1()
	ex2()
}

func ex2() {
	s := `[{"Name":"ku1"},{"Name":"ku2"}]`
	var users []user
	err := json.Unmarshal([]byte(s), &users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)

}

func ex1() {

	u1 := user{Name: "ku1"}
	u2 := user{Name: "ku2"}

	users := []user{u1, u2}

	bytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	fmt.Println("==============\n")
}
