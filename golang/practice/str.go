package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
	"strconv"
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
	// fmt.Println(strings.Replace("oink oink oink", "k", "ky", -1))
	asc()
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
	fmt.Println(strings.HasSuffix(str, "b"))
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
	// f := fibonacci(7)
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))
}


func asc() {
	// fmt.Println("=======================")
	// c:='a'
	// fmt.Println(c)
	// fmt.Println(string(c))
	// fmt.Println(string(97))
	// fmt.Println("=======================")
	// str1 := "97"
	// strInt,_ := strconv.Atoi(str1)
	// fmt.Println(strInt)
	// fmt.Println(string(strInt))
	// fmt.Println("=======================")
	// for _, char := range []rune("12.34") {
	// 	fmt.Println(char)//49
	// 	fmt.Println(string(char))//1
	// 	fmt.Println(string(49))
    // }
	// fmt.Println("=======================")
	// str2:="f3eee07574f0c9692854"
	// fmt.Println(str2[0:10])
	// fmt.Println("=======================")
	 i,_ := strconv.ParseInt("2E",16,10)
	 fmt.Println(string(i))
	 fmt.Println("=======================")
	 var s []string
	//  s = append(s, "f3","ee","e0","75","74")
	// 3031322E3334 
	// s = append(s,"30","31","32","2E","33","34")
	s = append(s,"30","31","32")
	 f := ascToFloat64(s)
	 fmt.Println(f)
	 fmt.Println("=====================sb==")
	 var sb []byte
	 for _,v := range s {
		 bs :=[]byte(v)
		 for _,i := range bs{
			sb = append(sb,i)
		}
	 }
	 fmt.Println("sb:",sb)
	 sbFloat ,_ :=strconv.ParseFloat(string(sb),64)
	 fmt.Println(sbFloat)
	}

func ascToFloat64(s []string ) float64 {
i := ""
for _,v := range s {
	m,_:= strconv.ParseInt(v, 16, 10)
	fmt.Println("m:", m)
	i = fmt.Sprintf("%s%s",i,string(m))
	fmt.Println("i", i)
}

f,_:= strconv.ParseFloat(i, 64)
fmt.Println("f",f)
return f
}
