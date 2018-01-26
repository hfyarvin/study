package main

import (
	"fmt"
	"time"
)

func test() {
	now := time.Now()
	// unix
	timestamp := now.Unix()
	timestampNano := now.UnixNano()
	fmt.Println(timestampNano)
	//timestamp to time
	newtime := time.Unix(timestamp, 0)

	//build time
	theDay := time.Date(newtime.Year(), newtime.Month()-1, newtime.Day(), 0, 0, 0, 0, newtime.Location())

	//utc
	fmt.Println(now.UTC())

	//add date
	nextDay := theDay.AddDate(0, 0, 1)
	fmt.Println(nextDay.Format("2006-01-02 15:04:05")) //Format Time
}

func main() {
	test()
}
