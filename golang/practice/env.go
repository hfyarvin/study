/**
 * 获取环境变量
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	gopath := os.Getenv("GOPATH")
	p := formatPathOfWindows(gopath)
	fmt.Println(gopath)
	fmt.Println(p)
}

func formatPathOfWindows(path string) string {
	pathArr := strings.Split(path, "\\")
	path = strings.Join(pathArr, "/")
	return path
}
