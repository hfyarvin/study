package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	q5()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

//-------------------------
type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu //指针只保留了最后一个student的内容
	}

	for k, v := range m {
		fmt.Println("k: ", k)
		fmt.Println("v: ", v)
	}
}

/**
 * sync包???
 */
func q1() {
	runtime.GOMAXPROCS(1) //直接设置环境变量$GOMAXPROCS改变程序使用的CPU数量
	wg := sync.WaitGroup{}
	wg.Add(20) //添加或者减少等待goroutine的数量
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("------first for----------")
			fmt.Println("i: ", i)
			wg.Done() //相当于Add(-1)
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("------second for----------")
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait() //执行阻塞，直到所有的WaitGroup数量变成0
}

//----------------------------------------------
var waitgroup sync.WaitGroup

func q3() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go sync1(i)
	}

	waitgroup.Wait()
	fmt.Println("donw!")
}

func sync1(showNum int) {
	fmt.Println(showNum)
	waitgroup.Done()
}

//----------------------------------------------
func q4() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//----------------------------------------------
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func q5() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
*/

//----------------------------------------------
