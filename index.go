package main

import (
	parser "github.com/MUR4SH/IOT_Practice2/tree/main/modules"

	"fmt"
)

func main() {
	_, _, h := parser.JSON_decode(parser.JSON_encode(1, "text1", "text2"))
	fmt.Println(h)
}
