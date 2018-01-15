/**
 * 打点器,定时做一些事情
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500) //每隔500毫秒
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop() //定时器结束
	fmt.Println("Ticker stopped")
}
