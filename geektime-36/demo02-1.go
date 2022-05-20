package main

//02丨命令源码文件
import (
	"flag"
	. "fmt"
	"os"
)

var name string

func init() {
	// 需在此处添加代码。[2]
	flag.StringVar(&name, "name1", "everyone", "The greeting object.")
}

// go run demo02-1.go -name1=cdb
func main() {
	//解析命令行传参到真正变量
	flag.Parse()
	flag.Usage = func() {
		Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	if name == "cdb" {
		println("果然是你红红火火恍恍惚惚或或或或或或或或或")
	}
	// 需在此处添加代码。[3]
	Printf("Hello, %s!\n", name)

}
