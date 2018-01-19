/**
 * 复制map不影响原map
 */
package main

import (
	// "encoding/json"
	"fmt"
)

var h map[string]string

func init() {
	h = map[string]string{
		"a": "aa",
		"b": "bb",
	}
}

func main() {
	// // 复制map不影响原map
	// bs, _ := json.Marshal(h)
	// var h2 map[string]string
	// err := json.Unmarshal(bs, &h2)
	// if err != nil {
	// 	fmt.Println("error: ", err.Error())
	// }

	// h2["c"] = "cc"

	// if _, ok := h["c"]; !ok {
	// 	fmt.Println("key is not exist!")
	// }

	// if _, ok := h2["c"]; ok {
	// 	fmt.Println("key is exist!")
	// }

	test()
}

func GetMap() map[string]string {
	var header map[string]string
	header = map[string]string{
		"a": "aa",
	}
	return header
}

func test() {
	a := GetMap()
	a["b"] = "bb"
	b := GetMap()
	if _, ok := a["b"]; ok {
		fmt.Println("key is xist!")
	}

	if _, ok := b["b"]; !ok {
		fmt.Println("key is not exist!")
	}
}
