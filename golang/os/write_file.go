/**
 *四种写文件的方式
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil" //io工具包
	"os"
)

/**
 *使用panic标注出所发生的错误
 **/
func check(e error) {
	if e != nil {
		panic(e) //panic 和recover
	}
}

/**
 *判断文件是否存在
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {
	var filename = "./output1.txt"
	var f *os.File
	var err error
	/***************************** 判断文件是否存在 ***********************************************************************/
	if checkFileIsExist(filename) {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //open file
		fmt.Println("file exist!")
	} else {
		f, err = os.Create(filename)
		fmt.Println("file not exist")
	}
	defer f.Close()
	check(err)
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	n, err := io.WriteString(f, "/io.WriteString\n") //write file(string),从文件尾续接
	check(err)
	fmt.Printf("write %d bytes\n", n)
	/***************************** 第二种方式: 使用 ioutil.WriteFile 写入文件 ********************** **********************/
	err = ioutil.WriteFile("./output2.txt", []byte("/ioutil.WriteFile\n"), 0666) //覆盖写入文件（字节数组）（相当于重新创建文件）
	check(err)
	/***************************** 第三种方式:  使用 File(Write,WriteString) 写入文件 *************************************/
	n2, err := f.Write([]byte("/File(Write)\n"))
	check(err)
	fmt.Printf("write %d bytes\n", n2)
	n3, err := f.WriteString("/File(WriteString)\n") //写入
	check(err)
	fmt.Printf("write %d bytes\n", n3)
	/***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 *********************************************/
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("/bufio.NewWriter\n")
	check(err)
	fmt.Printf("write %d bytes\n", n4)
	w.Flush() //bufio写入必须用到
}
