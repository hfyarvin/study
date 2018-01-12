/**
 * 100级，每次可以行走1级或者2级，一共有多少种方式
 */
package main

import (
	"fmt"
)

type M struct {
	Value int
}

var ml []M
var mmList [][]M

func main() {
	s := insertTest(0, mmList)
	fmt.Println(len(s))
	fmt.Println(s[1])
	fmt.Println(s[2])
}

func init() {
	for i := 0; i < 101; i++ {
		var m M
		m.Value = 1
		ml = append(ml, m)
	}

	mmList = insertStair(ml)
}

//再某种情况下，插入1个两级阶梯，有多少种情况,将之返回
func insertStair(mlist []M) [][]M {
	var list [][]M
	for i := 1; i < 100; i++ {
		if mlist[i].Value == 1 && mlist[i+1].Value == 1 {
			nm := copy(mlist)
			nm[i].Value = 0
			nm[i+1].Value = 0
			list = append(list, nm)
		}
	}
	return list
}

func copy(mlist []M) (l []M) {
	for _, item := range mlist {
		l = append(l, item)
	}
	return l
}

// 判断两种情况是否一样
func equal(m1, m2 []M) bool {
	e := true
	for i := 1; i < 101; i++ {
		if m1[i].Value != m2[i].Value {
			e = false
			break
		}
	}

	return e
}

func join(list1, list2 [][]M) [][]M {
	var newList [][]M
	newList = list1
	for _, item := range list2 {
		if !isContain(list1, item) { //如果不包含就将情况添加下去
			newList = append(newList, item)
		}
	}
	return newList
}

func isContain(list [][]M, ml []M) bool {
	is := false
	for _, item := range list {
		if equal(item, ml) {
			is = true
			break
		}
	}
	return is
}

//插入n次花多少次
func insertTest(n int, mm [][]M) [][]M {
	// newlist := mm
	if n == 0 {
		return mm
	}
	if n == 1 {
		var newlist [][]M
		for _, item := range mm {
			s := insertStair(item)
			newlist = join(s, newlist)
		}
		return newlist
	} else {
		ret := insertTest(n-1, mm)
		var newlist [][]M
		for _, item := range ret {
			s := insertStair(item)
			newlist = join(s, newlist)
		}
		return newlist
	}
}
