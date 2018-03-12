package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// file exist?
func checkFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}

	return exist
}

func main() {
	var writeString = "test n"
	var filename = "./output1.txt"
	var f *os.File
	var err1 error

	// 第一种方式: 使用 io.WriteString 写入文件
	if checkFileExist(filename) {
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666)
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename)
		fmt.Println("文件不存在")
	}

	check(err1)
	n, err1 := io.WriteString(f, writeString) //写入
	check(err1)
	fmt.Println("write %d bytes", n)

	// 第二种方式: 使用 ioutil.WriteFile 写入文件
	var d1 = []byte(writeString)
	err2 := ioutil.WriteFile("./output2.txt", d1, 0666)
	check(err2)

	//  第三种方式:  使用 File(Write,WriteString) 写入文件
	f, err3 := os.Create("./output3.txt")
	check(err3)
	defer f.Close()
	n2, err3 := f.Write(d1) //写入文件
	check(err3)
	fmt.Printf("写入 %d 个字节", n2)
	n3, err3 := f.WriteString("writesn")
	fmt.Printf("写入 %d 个字节n", n3)
	f.Sync()

	// 第四种方式:  使用 bufio.NewWriter 写入文件
	w := bufio.NewWriter(f)
	n4, err3 := w.WriteString("bufferedn")
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}
