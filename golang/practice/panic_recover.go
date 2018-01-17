package main

import "fmt"

func main() {
	/*
	   先defer的后执行
	   recover后输出panic中的信息
	*/

	defer func() {
		if err := recover(); err != nil {

			fmt.Print(err)
		} else {
			fmt.Print("no")
		}

	}()
	defer func() {
		panic("1111111111111")
	}()
	panic("22222222222")

}
