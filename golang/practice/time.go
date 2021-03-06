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
	test2()
}

//时间与字符串
func test2(){
	//获取本地location
	toBeCharge := "2015-01-01 01:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600
	fmt.Println(theTime.Hour()+1)

	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)  
}
