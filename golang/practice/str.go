package main

import (
	"bytes"
	"fmt"
	"strings"
)

/**
 * 拼接字符串
 */
var (
	h = "hello"
	w = "world"
)

func way1() {
	s := h + "," + w
	fmt.Println(s)
}

func way2() {
	s := fmt.Sprintf("%s,%s", h, w)
	fmt.Println(s)
}

func way3() {
	s := strings.Join([]string{h, w}, ",")
	fmt.Println(s)
}

func way4() {
	var buffer bytes.Buffer
	buffer.WriteString(h)
	buffer.WriteString(",")
	buffer.WriteString(w)
	s := buffer.String()
	fmt.Println(s)
}

func main() {
	way4()
}
