package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
test2()
}

func test1()  {
	var v interface{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		v = i

		if (r.Intn(100) % 2) == 0 {
			v = "hello"
		}

		if _,ok := v.(int); ok { //如果类型ok
			fmt.Printf("%d\n", v)
		}
	}
}


func test2()  {
    add(1, 2)
    add(int16(1), int16(2))
    add(float32(1.1), float32(2.2))
    add(float64(1.1), float64(2.2))
    add(true, false)
}

func add(a, b interface{})  {
	switch t := a.(type) {
	case int:
		fmt.Printf("type [%T] add res[%d]\n", t, a.(int)+b.(int))
	case int16:
		fmt.Printf("type [%T] add res[%d]\n", t, a.(int16)+b.(int16))
	case float32:
		fmt.Printf("type [%T] add res[%d]\n", t, a.(float32)+b.(float32))
	case float64:
		fmt.Printf("type [%T] add res[%d]\n", t, a.(float64)+b.(float64))
	default:		
		fmt.Printf("type [%T] not support!\n", t)
	}
}