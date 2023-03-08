package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrMath = errors.New("math: sqrt root of negative number")

func main() {
	fmt.Printf("%T\n", ErrMath)
	_, err := sqrt2(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt2(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMath
	}

	return 42, nil
}
