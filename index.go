package main

import (
	"fmt"

	parser "~/Users/alex-murashev/Documents/GitHub/IOT_Practice2/modules/coderEncoder.go"
)

func main() {
	_, _, h := parser.JSON_decode(parser.JSON_encode(1, "text1", "text2"))
	fmt.Println(h)
}
