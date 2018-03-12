package main

import (
	"fmt"
)

func main() {
	t1()
}

func t1() {
	c1 := make(chan int, 1)
	c1 <- 10
	v, ok := <-c1
	if ok {
		fmt.Println("v: ", v)
	}
	close(c1) //关闭
	if _, ok := <-c1; !ok {
		fmt.Println("no value")
	}
}
