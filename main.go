package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	runes := []rune(text)
	res := string(runes[0:width])

	fmt.Printf("%v", res)
}
