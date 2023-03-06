package main

import (
	"fmt"
	"io"
	"os"
)

/*
Encoder : Writer, marshar와 비슷

	// 유저가 쓰기

Decoder: Reader, unmarshar와 비슷

	// 유저가 읽기
*/
func main() {
	fmt.Fprintln(os.Stdout, "Hello std out")
	io.WriteString(os.Stdout, "Hello Write String")

}
