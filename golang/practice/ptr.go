package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {

	m := pase_map()

	for k, v := range m {
		fmt.Printf("key = %s,value =%v\n", k, v)
	}

}

func pase_map() map[string]*student {

	m := make(map[string]*student)

	stu := []student{{"joy", 12}, {"lei", 14}}

	for _, v := range stu { //range只保留最后一次遍历v的地址,所以也只会保留最后一次遍历的值//range只保留最后一次遍历v的地址
		fmt.Printf("addr: %p\n", &v)
		m[v.Name] = &v
	}
	return m
}
