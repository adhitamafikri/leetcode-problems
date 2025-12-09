package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	trimmed := strings.Trim(s, " ")
	length := len(trimmed)
	result := 0

	for i := length - 1; i >= 0; i-- {
		// increment as long as the current character ascii is not an ascii of whitespace
		if trimmed[i] != 32 {
			result++
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println("Length of last word")
	fmt.Println(lengthOfLastWord(("Hello World")))
	fmt.Println(lengthOfLastWord(("   fly me   to   the moon  ")))
	fmt.Println(lengthOfLastWord(("luffy is still joyboy")))
	fmt.Println(lengthOfLastWord(("a")))
}
