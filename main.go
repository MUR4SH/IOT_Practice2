package main

import (
	"fmt"
)

func main() {
	_, t, _ := HTML_decode(HTML_encode(1, "text1", "text2"))
	fmt.Println(t)
}
