package main 

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println("end")
	quit <- 0
}

func main()  {
	runtime.GOMAXPROCS(2) // 最多使用2个核

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<- quit
	}
}