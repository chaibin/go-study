package main

import "fmt"

func main() {
	var a, b string
	a = "aaa"
	b = "bbbb"
	a, b = Func0(a, b)
	fmt.Println(a, b)
}
func Func0(a string, b string) (string, string) {
	a += "---"
	b += "===="
	return a, b
}
