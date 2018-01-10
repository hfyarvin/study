
/*
给定一个候选数字数组 (C) (不含重复项) 和目标数字 (T)，在 C 中找到所有唯一的组合，使得组合中的数之和为 T。

C 中的同一个数字可以无限制重复使用。

注意：

所有数字，包括目标数字，都是正数
结果数组中不能包含重复的组合
例如：给定一个候选数组： [2, 3, 6, 7] 和 目标数字 7,
*/
package main

import "fmt"

func main() {
	C := []int{2,3,6,7}
	targetNum := 7
}

type Ret struct{
	Result [][]int
}

func GetResultArr(arr []int, target int) []int{
	loopTimes := target / GetResultArr(arr)
	empty := new([]int)
	ret := new(Ret)
	for _,i := range arr {
		resArr := []int
		sum := j
		resArr = append(resArr, i)
		for t := 0; t < loopTimes -1 ; t++  {
			for _,j := range arr {
				if j > target || (sum + j) > target {
					continue
				} else {
								
				}
			}
		}
	}
}

func GetMin(arr []int) int {
	min := arr[0]
	for _,i := range arr {
		if i < min {
			min = i
		}
	}

	return min
}
