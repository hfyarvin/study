package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用缓冲scanner来包裹无缓冲的`os.Stdin`可以让我们
	// 方便地使用`Scan`方法，这个方法会将scanner定位到下
	// 一行的位置
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
