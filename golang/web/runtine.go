//goroutine协程
//channel
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 4; i++ {
		/*runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		  默认情况下，在Go 1.5将标识并发系统线程个数的runtime.GOMAXPROCS的初始值由1改为了运行环境的CPU核数
		*/
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	bufferChan()
}

func routine_test() {
	go say("arvin")
	say("I'm")
}

//channel
/*
ch <- v 发送v到channel ch
v := <- ch 从ch中接收数据
值：先进后出
ch := make(chan type, value)
value=0,channel 是无缓冲阻塞读写的
value>0,有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入
*/
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func channel_test() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}

//缓存型channel
func bufferChan() {
	c := make(chan int, 1)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

//读写channel
func chan_rw() {
	c := make(chan int, 10)
	// 读
	for i := range c {

	}
	//关
	close(c)
	//检测是否关闭
	v, ok := <-ch
}
