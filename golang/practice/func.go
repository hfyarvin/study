package main

import (
	"fmt"
)

/*
1. niladic function: 没有参数的函数通常被称为 niladic 函数
2. call by value
3. call by reference
*/

func main() {
	// funcExample1(1, 2, 3)
	// b := B(20)
	// fmt.Println(b)
	Ptr()
}

//传递变长参数(可以是同类型，也可以不是)
func funcExample1(a, b int, l ...int) {
	fmt.Println(a)
	fmt.Println(b)
	for _, v := range l {
		fmt.Println("the arg :", v)
	}
}

func B(a int) int {
	r := 1
	for i := 0; i < a; i++ {
		r *= 2
	}
	return r
}

//以函数当作类型
type binOp func(int, int) int

//传递指针
func Ptr() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
