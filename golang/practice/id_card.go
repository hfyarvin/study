package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	idCard := "511025199410280934"
	b := VerifyIdCard(idCard)
	fmt.Println(b)
}

//获取age字符串
func GetAgeStr(idCard string) string {
	if idCard == "" || len(idCard) < 18 {
		return ""
	}

	idCardArr := strings.Split(idCard, "")
	birthday := fmt.Sprintf("%s-%s-%s", strings.Join(idCardArr[6:10], ""), strings.Join(idCardArr[10:12], ""), strings.Join(idCardArr[12:14], ""))
	return birthday
}

//根据年龄字符串获取年龄(虚岁)
func GetAge(ageStr string) int {
	if ageStr == "" {
		return 0
	}

	arr := strings.Split(ageStr, "-")
	birthYear, _ := strconv.Atoi(arr[0])
	birthMonth, _ := strconv.Atoi(arr[1])
	birthDay, _ := strconv.Atoi(arr[2])
	y, m, d := time.Now().Date()
	age := 0

	if (birthYear > y) || (birthYear == y && birthMonth > int(m)) || (birthMonth == int(m) && birthDay > d) {
		return 0
	} else {
		yearDiffer := y - birthYear
		age = yearDiffer
		if (birthMonth < int(m)) || (birthMonth == int(m) && birthDay <= d) {
			age += 1
		}
	}
	return age
}

func test() {
	y, m, d := time.Now().Date()
	fmt.Println(y, int(m), d)
	return
}

// 将身份证号转为数字数组--最后一位会有误
func IdCardToNumArray(idCard string) []int {
	s := []byte(idCard)
	var arr []int
	for _, v := range s {
		t, _ := strconv.Atoi(string(v))
		arr = append(arr, t)
	}
	if len(arr) != 18 {
		return nil
	}
	return arr
}

func GetCheckNum(idCardNumArr []int) int {
	var res int
	var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	for i := 0; i < 17; i++ {
		res += idCardNumArr[i] * wi[i]
	}
	return (res % 11)
}

func VerifyIdCard(idCard string) bool {
	v := false
	idCardNumArr := IdCardToNumArray(idCard)
	if idCardNumArr == nil {
		return false
	}
	checkNum := GetCheckNum(idCardNumArr)
	a18 := [11]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2}
	var checkByte byte
	for i := 0; i < 11; i++ {
		if i == checkNum {
			checkByte = a18[i]
			break
		}
	}

	idCardBytes := []byte(idCard)
	t, _ := strconv.Atoi(string(idCardBytes[17]))
	if checkNum == 2 && checkByte == idCardBytes[17] {
		v = true
	}
	if checkByte == byte(t) {
		v = true
	}
	return v
}
