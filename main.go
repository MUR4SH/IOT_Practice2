package main

import (
	"fmt"
)

func main() {
	_, _, h := JSON_decode(JSON_encode(1, "text1", "text2"))
	fmt.Println(h)
}
