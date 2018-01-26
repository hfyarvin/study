package main

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

func main() {
	Test()
}

func Test() {
	c1 := GenerateRandomCode()
	fmt.Println(c1)
	//大写
	fmt.Println(strings.ToUpper(c1))
}

//生成8位数随机码
func GenerateRandomCode() string {
	str, _ := uuid.NewV4()
	u := str.String()
	c := strings.Split(u, "-")
	return c[0]
}
