package main

import (
	"fmt"
	"os"
)

func main() {
	var s1, s2 string
	for i := 1; i < len(os.Args); i++ {
		s1 += s2 + os.Args[i]
		s2 = " "
	}
	fmt.Println(s1)
	i := 10
	for i > 0 {
		i--
		fmt.Println(i)
	}
}
