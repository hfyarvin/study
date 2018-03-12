package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	t1()
}

func incr(p *int) int {
	*p++
	return *p
}

func t1() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
	//工作目录
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("os.Getwd failed:  ", err)
	}
	log.Println("Getwd: ", cwd)
}
