/**
 * sha1加密
 */
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "arvin.Wong"
	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(bs)
}
