package main

import (
	"fmt"
)

func main() {
	// var pointer *int
	pointer := new(int)
	*pointer = 3
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(&pointer)

	p2 := new(*int)
	p2 = &pointer
	fmt.Println(p2)
	fmt.Println(*p2)
}
