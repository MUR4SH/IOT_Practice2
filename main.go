package main

import (
	"fmt"
)

func main() {
	i, _, _ := HTML_decode(HTML_encode(1, "text1", "text2"))
	fmt.Println(i)
}
