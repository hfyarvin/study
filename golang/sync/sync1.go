package main

import (
	"fmt"
)

/**
* var ch = make(chan int) //无缓冲channel,等同于make(chan int ,0),无缓冲channel在读和写是都会阻塞
* var ch = make(chan int,10) //有缓冲channel,缓冲大小是５,有缓冲channel在向channel中存入数据没有达到channel缓存总数时，可以一直向里面存
 */
func Afunction(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {
	ch := make(chan int) //无缓冲channel
	go Afunction(ch)
	fmt.Println("input:")
	ch <- 1
}

/**
func main() {
	ch := make(chan int) //无缓冲的channel
	//只是把这两行的代码顺序对调一下
	ch <- 1
	go Afuntion(ch)

	// 输出结果：
	// 死锁，无结果
}
*/
