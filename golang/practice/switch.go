package main

import (
	"fmt"
	"time"
)

func main() {
	switchExample1(100)
	switchExample2(2)
	switchExample2(22)
	switchExample3(6) //fallthrough继续执行之后的代码
	season := time.Now().Month()
	fmt.Println("The season of now is", GetSeasonBySeasonNUm(int(season)))
}

func switchExample1(num1 int) {

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}

func switchExample2(num1 int) {
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}

//fallthrough:继续执行之后的代码
func switchExample3(num1 int) {
	switch num1 {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func GetSeasonBySeasonNUm(seasonNum int) string {
	switch seasonNum {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	}
	return "Season unknown"
}
