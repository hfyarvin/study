/**
 * 排序算法
 */
package main

import (
	"fmt"
)

/**************** 冒泡排序 *********************/
func ssort(a []int) []int {
	times := 0
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				times++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(times)
	return a
}

/**************** 鸡尾酒排序 *********************/
func CocktailSort(a []int) []int {
	times := 0
	left := 0
	right := len(a) - 1
	for left < right {
		for i := left; i < right; i++ {
			if a[i] > a[i+1] {
				times++
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		fmt.Println(left, right)
		right--
		for i := right; i > left; i-- {
			if a[i-1] > a[i] {
				times++
				a[i-1], a[i] = a[i], a[i-1]
			}
		}
		left++
	}
	fmt.Println(times)
	return a
}

/**************** 选择排序 *********************/
func SelectionSort(a []int) []int {
	times := 0
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ { //找到未排序的最小
			if a[j] < a[min] {
				times++
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
	fmt.Println(times)
	return a
}

/**************** 插入排序 *********************/
func InsertionSort(a []int) []int {
	//
	times := 0
	for i := 0; i < len(a); i++ {
		get := a[i]
		j := i - 1
		for j >= 0 && a[j] > get {
			times++
			a[j+1] = a[j] //如果该手牌比抓到的牌大，就将其右移
			j--           //再和更左边一张对比
		}
		a[j+1] = get
	}
	fmt.Println(times)
	return a
}

/**************** 插入排序(二分法) *********************/
func InsertionSortDichotomy(a []int) []int {
	times := 0
	for i := 0; i < len(a); i++ {
		get := a[i]
		left := 0
		right := i - 1

		for left <= right {
			mid := (left + right) / 2
			if a[mid] > get {
				times++
				right = mid - 1
			} else {
				times++
				left = mid + 1
			}
		} //插入位置（left）
		for j := i - 1; j >= left; j-- {
			times++
			a[j+1] = a[j]
		}
		a[left] = get
	}
	fmt.Println(times)
	return a
}

/**************** 希尔排序 *********************/
func ShellSort(a []int) []int {
	times := 0
	h := 0
	for h <= len(a) {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < len(a); i++ {
			j := i - h
			get := a[i]
			for j >= 0 && a[j] > get {
				times++
				a[j+h] = a[j]
				j = j - h
			}
			a[j+h] = get
		}
		h = (h - 1) / 3
	}
	fmt.Println(times)
	return a
}
func main() {
	a := []int{5, 2, 9, 4, 7, 6, 1, 3, 8}
	b := ShellSort(a)
	fmt.Println(b)
}
