package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

}
