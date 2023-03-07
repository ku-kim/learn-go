package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("1234.txt")
	if err != nil {
		//fmt.Println(err)
		//log.Println(err)
		//log.Fatalln(err)
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called. defered functions don't run")
}
