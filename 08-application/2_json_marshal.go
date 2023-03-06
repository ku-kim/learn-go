package main

import (
	"encoding/json"
	"fmt"
)

/*
마샬링 (e.g. type -> json)
 소프트웨어는 바이트 단위로 데이터를 인식하기때문에,
바이트를 정수로 읽어들일지,문자로 읽어들일지에 따라 출력되는 값이 달라진다.
예를들어, 메모리에 올려진 100을 정수로 인식하면 그대로 100이 되는것이고,
이것을 문자로 읽으면 'd'라는 문자가 된다.
메모리에 올려진 값을 정수 또는 다른 타입으로 변환하는 과정을 마샬링(Marshaling)이라고 한다.


반대로 언마샬링은
바이트 슬라이스나 문자열을 논리적 자료구조로 변경하는 것( json -> type)
*/

func main() {
	// Marshal은 소문자로 field가 셋업되면 인식하지 못함
	type person struct {
		First string
		Last  string
		Age   int
	}

	p1 := person{First: "James",
		Last: "Bond",
		Age:  26,
	}

	p2 := person{First: "James2",
		Last: "Bond2",
		Age:  26,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
