/**
* 斐波那契数列
 */
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("f(%d): %d", i, fibonacci(i)))
	}
}

func fibonacci(n int) int {
	if n >= 3 {
		return fibonacci(n-1) + fibonacci(n-2)
	} else if n == 1 || n == 2 {
		return 1
	}
	return 0
}
