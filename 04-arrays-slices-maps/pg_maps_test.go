package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapsOrder(t *testing.T) {
	m := map[string]int{
		"a": 1,
	}
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8

	o1 := make([]string, 5, 5)
	for k, _ := range m { // 순서 보장하지 않음
		o1 = append(o1, k)
	}

	o2 := make([]string, 5, 5)
	for k, _ := range m { // 순서 보장하지 않음
		o2 = append(o2, k)
	}

	if reflect.DeepEqual(o1, o2) != true {
		fmt.Println(o1)
		fmt.Println(o2)
		fmt.Println("확률적으로 map은 순서 보장하지 않음(o1, o2는 동등하지 않음)")
	} else {
		fmt.Println(o1)
		fmt.Println(o2)
		t.Error("확률적으로 같으거니 테스트 재시도 할 것")
	}

}

func TestMapValueInListOrder(t *testing.T) {
	m := map[string][]int{
		"a": {-1},
		"b": {-1},
	}

	for i := 0; i < 10; i++ {
		m["a"] = append(m["a"], i)
	}

	for i := 0; i < 3; i++ {
		m["b"] = append(m["b"], i)
	}

	fmt.Println(m["a"])
	fmt.Println(m["b"])

	for k, v := range m {
		fmt.Println(k, v) // a, b 순서는 map(key)에 따라 랜덤이겠지만 해당 value는 list이기에 순서 지키고 있음
	}

}
