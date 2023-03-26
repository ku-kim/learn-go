package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

type Person struct {
	Name     string
	OrderNo1 int
	OrderNo2 int
}

type Person2 struct {
	Name     string
	OrderNo1 int
	OrderNo2 int
	StartAt  time.Time
}

func baseTime(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Seoul")
	year, month, day := t.In(loc).Date()
	onTime := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return onTime
}

// https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters
func TestStructSort(t *testing.T) {
	person1 := Person{Name: "b", OrderNo1: 3, OrderNo2: 0}
	person2 := Person{Name: "b", OrderNo1: 2, OrderNo2: 2}
	person3 := Person{Name: "b", OrderNo1: -1, OrderNo2: 2}
	person4 := Person{Name: "b", OrderNo1: 10, OrderNo2: 2}
	person5 := Person{Name: "b", OrderNo1: 2, OrderNo2: 1}
	person6 := Person{Name: "a", OrderNo1: 1, OrderNo2: 2}

	peopleList := []Person{}

	peopleList = append(peopleList, person1)
	peopleList = append(peopleList, person2)
	peopleList = append(peopleList, person3)
	peopleList = append(peopleList, person4)
	peopleList = append(peopleList, person5)
	peopleList = append(peopleList, person6)

	fmt.Println(peopleList)
	sort.Slice(peopleList, func(i, j int) bool {
		if peopleList[i].Name != peopleList[j].Name {
			return peopleList[i].Name < peopleList[j].Name
		} else if peopleList[i].OrderNo1 != peopleList[j].OrderNo1 {
			return peopleList[i].OrderNo1 < peopleList[j].OrderNo1
		}
		return peopleList[i].OrderNo2 < peopleList[j].OrderNo2
	})

	fmt.Println(peopleList)
}

func TestStructSort2(t *testing.T) {
	today := baseTime(time.Now())
	person1 := Person2{Name: "b", OrderNo1: 3, OrderNo2: 0, StartAt: today.AddDate(0, 0, 2)}
	person2 := Person2{Name: "b", OrderNo1: 2, OrderNo2: 2, StartAt: today.AddDate(0, 0, 1)}
	person3 := Person2{Name: "b", OrderNo1: -1, OrderNo2: 2, StartAt: today.AddDate(0, 0, 0)}
	person34 := Person2{Name: "b", OrderNo1: -1, OrderNo2: 0, StartAt: today.AddDate(0, 0, 0)}
	person4 := Person2{Name: "b", OrderNo1: 10, OrderNo2: 2, StartAt: today.AddDate(0, 0, 0)}
	person5 := Person2{Name: "a", OrderNo1: 2, OrderNo2: 1, StartAt: today.AddDate(0, 0, 0)}
	person6 := Person2{Name: "a", OrderNo1: 1, OrderNo2: 2, StartAt: today.AddDate(0, 0, 0)}

	peopleList := []Person2{}

	peopleList = append(peopleList, person1)
	peopleList = append(peopleList, person2)
	peopleList = append(peopleList, person3)
	peopleList = append(peopleList, person34)
	peopleList = append(peopleList, person4)
	peopleList = append(peopleList, person5)
	peopleList = append(peopleList, person6)

	fmt.Println(peopleList)
	sort.Slice(peopleList, func(i, j int) bool {
		iv, jv := peopleList[i], peopleList[j]
		switch {
		case iv.StartAt != jv.StartAt:
			return iv.StartAt.Before(jv.StartAt)
		case iv.Name != jv.Name:
			return iv.Name < jv.Name
		case iv.OrderNo2 != jv.OrderNo2:
			return iv.OrderNo2 < jv.OrderNo2
		default:
			return iv.OrderNo1 < jv.OrderNo1
		}
	})

	fmt.Println(peopleList)
}
