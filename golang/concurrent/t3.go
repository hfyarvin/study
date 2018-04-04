package main 

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine

		fmt.Printf("%d", i)
	}

	quit <- 0
}

func main()  {
	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<- quit
	}
}
/**
 *主动让出CPU时间的方式仍然是在单核里跑。但手工地切换goroutine导致了看上去的“并行”。
 *关于runtime包几个函数： 
 *Gosched() //让出cpu 
 *NumCPU()//返回当前系统的CPU核数量 
 *GOMAXPROCS() //设置最大的可同时使用的CPU核数 
 *Goexit() //退出当前goroutine(但是defer语句会照常执行)
*/ 