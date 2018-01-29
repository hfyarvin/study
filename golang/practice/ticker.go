/**
 * 打点器,定时做一些事情
 */
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ticker2()
}

func ticker1() {
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

func ticker2() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, os.Kill)

	go func() {
		sig := <-sigs
		switch sig {
		case os.Interrupt:
			fmt.Println("singal: Interrupt")
		case os.Kill:
			fmt.Println("signal: Kill")
		default:
			fmt.Println("signal: Options")
		}
		done <- true
	}()

	fmt.Println("awaiting signal")

	go JobTicker(done)
	<-done
	close(done)
	fmt.Println("exiting")
}

func JobTicker(done <-chan bool) {
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			fmt.Println("job...")
		}
	}
}
