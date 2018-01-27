package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	GetRandInt64ByRange(1, 7)
}

func RandInt() int {
	a := rand.Int()
	fmt.Printf("%d\n", a)
	return a
}

//[0,n)之间伪随机数
func RandIntn(n int) int {
	num := rand.Intn(n)
	fmt.Println(num)
	return num
}

func RandFloat32() float32 {
	timeNs := time.Now().UnixNano()
	rand.Seed(timeNs)
	a := rand.Float32()
	fmt.Printf("%2.2f\t", a) //%2.2f会自动四舍五入
	b := a * 100
	fmt.Printf("%2.2f\n", b)
	return b
}

//产生随机数
func GetRandIntn(n int) int {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(n)
	fmt.Println(x)
	return x
}

//随机生成一个范围的整数,包含max
func GetRandInt64ByRange(min, max int64) int64 {
	var a int64
	a = min
	a += rand.New(rand.NewSource(time.Now().Unix())).Int63n(max + 1 - min)
	fmt.Println(a)
	return a
}
