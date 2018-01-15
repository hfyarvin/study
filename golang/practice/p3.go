/**
 * p2.go 问题
 * 100级，每次可以行走1级或者2级，一共有多少种方式
 */
package main

import (
	"fmt"
)

func main() {
	// for i := uint64(97); i <= 100; i++ {
	// 	a := Sum(i)
	// 	b := Sum(i - 1)
	// 	c := Sum(i - 2)
	// 	fmt.Printf("%d:%d\n", i, a)
	// 	fmt.Printf("%d:%d\n", i-1, b)
	// 	fmt.Printf("%d:%d\n", i-2, c)
	// 	if a != b+c {
	// 		fmt.Printf("Sum %d is the limit", i)
	// 	}
	// }
}

func Sum(n uint64) uint64 {
	if n > 6 {
		return 8*Sum(n-5) + 5*Sum(n-6)
	}

	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else { //n>=2
		return Sum(n-1) + Sum(n-2)
	}
}

/**
 * 菲波那契数列：F(n) = F(n-1) + F(n-2)
 **/
