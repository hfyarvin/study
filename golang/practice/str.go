package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
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
	str := "arvin"
	fmt.Println(string(str[0]))
}

func stringsTest() {
	str := "a/b/c"
	//切割
	strArr := strings.Split(str, "/")
	fmt.Println(strArr)
	//大小写
	str = "a3B45Cdd"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	//长度len
	fmt.Println(len(str))
	//拼接字符串
	str += "aaa"
	fmt.Println(str)
	//第一个字节
	fmt.Println(str[0])
	fmt.Println(string(str[0]))
	//前后缀
	fmt.Println(strings.HasPrefix(str, "a"))
	fmt.Println(strings.HasSuffix(str, "b"))SelectionSort(a)
	// 包含
	fmt.Println(strings.Contains(str, "aaa"))
	//在字符串中的位置-1表示不包含
	fmt.Println(strings.Index(str, "a3"))
	fmt.Println(strings.LastIndex(str, "a3")) //最后一次出现
	//替换-1为全部替换(返回新字符串)
	str2 := strings.Replace(str, "a3", "b2", -1)
	fmt.Println(str2)
	//统计出现次数
	fmt.Println(strings.Count(str, "a"))
	//新建一个重复内容的字符串
	fmt.Println(strings.Repeat("a", 3))
	//utf8
	countCharacter()
}

func countCharacter() {
	// count number of characters:
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	f := fibonacci(7)
	f := fibonacci(7)
	f := fibonacci(7)
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))
}
