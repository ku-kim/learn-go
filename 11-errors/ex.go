package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type customErr2 struct {
	Msg  string
	Code int
}

func (c customErr2) Error() string {
	return fmt.Sprintf("error: %v, %v", c.Msg, c.Code)
}

func main() {

	//ex1()
	//ex2()
	ex3()
}

func ex3() {
	p := person{
		First:   "james",
		Last:    "bond",
		Sayings: []string{"haha", "11", "22"},
	}

	bs, err := toJson2(p)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(bs))
}

func toJson2(p person) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return []byte{}, customErr2{Msg: "haha", Code: 1234}
	}
	return bs, nil
}

func ex2() {
	p := person{
		First:   "james",
		Last:    "bond",
		Sayings: []string{"haha", "11", "22"},
	}

	bs, err := toJson(p)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(bs))
}

func toJson(p person) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return []byte{}, fmt.Errorf("Error:  %v", err)
		//return []byte{}, errors.New("Error:  haha")
	}

	return bs, nil
}

func ex1() {
	p := person{
		First:   "james",
		Last:    "bond",
		Sayings: []string{"haha", "11", "22"},
	}

	bs, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
