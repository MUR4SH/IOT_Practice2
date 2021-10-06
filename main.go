package main

import (
	"fmt"
)

func main() {
	_, _, h := XML_decode(XML_encode(1, "text1", "text2"))
	fmt.Println(h)
}
