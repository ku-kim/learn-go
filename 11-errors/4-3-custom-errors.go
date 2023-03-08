package main

import (
	"fmt"
	"log"
)

type CustomError struct {
	message string
	code    string
}

func (err CustomError) Error() string {
	return fmt.Sprintf("error: %v, %v", err.message, err.code)
}

func main() {
	_, err := sqrt4(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt4(f float64) (float64, error) {
	if f < 0 {
		return 0, CustomError{"커스텀에러야", "400"}
	}

	return 42, nil
}
