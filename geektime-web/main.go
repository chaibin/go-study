package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	println("hello go")
	//utf8 汉字3-4个字节 len是计算字节
	println(len("我长度是14"))
	println(utf8.RuneCountInString("我长度是6哇"))
	fmt.Println()

}
