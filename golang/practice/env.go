/**
 * 获取环境变量
 */

package main

import (
	"fmt"
	"os"
	"strings"

	"runtime"
)

func main() {
	OsEnv()
	RuntimeEnv()
}

//格式化路径一下下
func formatPathOfWindows(path string) string {
	pathArr := strings.Split(path, "\\")
	path = strings.Join(pathArr, "/")
	return path
}

func OsEnv() {
	gopath := os.Getenv("GOPATH")
	p := formatPathOfWindows(gopath)
	fmt.Println(gopath)
	fmt.Println(p)
	dir, _ := os.Getwd()
	fmt.Println("dir: ", dir)
}

func RuntimeEnv() {
	fmt.Println(runtime.GOOS) //操作系统
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOARCH)
}
