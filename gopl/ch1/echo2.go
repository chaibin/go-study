package main

import (
	"fmt"
	"os"
)

func main() {
	s1, s2 := "", ""

	for _, arg := range os.Args[1:] {
		s1 += s2 + arg
		s2 = " "
	}
	fmt.Println(s1)
	testVar()
}

func testVar() {
	var s1 string
	s1 = "sss"
	fmt.Println(s1)
}
