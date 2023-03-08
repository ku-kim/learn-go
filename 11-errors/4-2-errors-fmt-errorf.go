package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt3(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt3(f float64) (float64, error) {
	if f < 0 {
		// fmt.Errorf, errors.New() 리턴하는데 format 형식으로 입력받아 문자열 조합 가능
		//변수로 빼도 되지만 errors,New()는 전역변수 가능하지만 여기는 매개변수 받아서 단점
		//errVar := fmt.Errorf("math: sqrt root of negative number, %f", f)
		return 0, fmt.Errorf("math: sqrt root of negative number, %f", f)
	}

	return 42, nil
}
