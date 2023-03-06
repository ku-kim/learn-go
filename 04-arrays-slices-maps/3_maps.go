package main

import "fmt"

func main() {
	basic_map()

	map_add_elements()
	map_delete_elements()
}

func map_delete_elements() {
	m := map[string]int{
		"ku": 1,
	}
	m["haha"] = 2 // map 데이터 추가

	delete(m, "haha")
	fmt.Println(m)
}

func map_add_elements() {
	m := map[string]int{
		"ku": 1,
	}
	m["haha"] = 2 // map 데이터 추가

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func basic_map() {
	m := map[string]int{
		"James": 1,
		"kukim": 2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	v, ok := m["what"] // 없으면 0, false
	fmt.Println(v, ok)

	if _, ok := m["what"]; ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}
