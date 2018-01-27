package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(IntFromFloat64(5.66))
}

//int to unint8
func UintFromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%s is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x) //整数部分和小数部分
		fmt.Println(whole, fraction)
		s := fmt.Sprintf("%2.2f", fraction)
		fmt.Println(s)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
