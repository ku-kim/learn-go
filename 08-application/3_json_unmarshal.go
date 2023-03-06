package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type person struct {
		First string `json:"FirstName"` // mapping 됨
		Last  string
		Age   int
	}

	// 뒤의 First는 FirstName과 매핑되지 않아서 들어가지 않음
	s := `[{"FirstName":"James","Last":"Bond","Age":26},{"First":"James2","Last":"Bond2","Age":26}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs) // []bytes는 uint8 로 type 되어있음

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
