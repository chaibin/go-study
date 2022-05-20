package main

//02丨命令源码文件
import (
	"flag"
	. "fmt"
)

var name1 string

func init() {
	// 需在此处添加代码。[2]
	flag.StringVar(&name1, "name1", "everyone", "The greeting object.")
}

// go run demo02.go -name1=cdb
func main() {
	//解析命令行传参到真正变量
	flag.Parse()
	// 需在此处添加代码。[3]
	Printf("Hello, %s!\n", name1)

	var i = 1

	var test = func(int2 *int) *int {
		int2 = &i
		return int2
	}
	testResult := test(&i)
	println(test(&i), testResult)
}
