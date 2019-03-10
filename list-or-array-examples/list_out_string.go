package main

import (
	"fmt"
	"strings"
)

func longestWord(s string) int {
	length := 0
	for _, word := range strings.Split(s, " ") {
		if len(word)%2 == 0 {
			length = len(word)
		}
		if length > 0 && length == (len(word)) {
			fmt.Println(word)
		}
	}
	return length
}

func main() {
	args := "Monday Tuesday Friday Sunday Wednesday abcdef"
	fmt.Printf("%v\n", longestWord(args))
}
