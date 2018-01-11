package main

import (
	"flag"
	"fmt"
)

var InputStrName = flag.String("name", "arvin", "input ur name!") //设定了默认值，返回一个xiangying类型指针
var InputAge = flag.Int("age", 20, "input ur age!")               //设定了默认值
var InputFlagvar int

func Init() {
	flag.IntVar(&InputFlagvar, "flagname", 1234, "help message for flagname.") //将flag绑定到一个变量
}

func main() {
	Init()

	flag.Parse() //解析

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg()) //获取参数数组与参数数量

	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name=", *InputStrName)
	fmt.Println("age=", *InputAge)
	fmt.Println("flagname=", InputFlagvar)
}
