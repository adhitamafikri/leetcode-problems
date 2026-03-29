package main

import (
	"fmt"
	"slices"
)

func reverseVowels(s string) string {
	// The idea:
	// 1. Have 2 pointers (p1 at 0, p2 at len - 1)
	// 2. Pointer points at vowel should stop, pointer at non-vowel should move
	// 3. Both pointer points to a vowel, do swap
	// 4. Stop the loop if p1 => p2

	// Caveats:
	// In Go, strings are immutable
	// Try to use conversion to byte

	p1, p2 := 0, len(s)-1
	b := []byte(s)
	v1, v2 := false, false

	for {
		if isVowel(b[p1]) {
			v1 = true
		} else {
			p1++
		}

		if isVowel(b[p2]) {
			v2 = true
		} else {
			p2--
		}

		if v1 && v2 {
			temp := b[p1]
			b[p1] = b[p2]
			b[p2] = temp

			p1++
			p2--
		}

		v1 = false
		v2 = false

		if p1 >= p2 {
			break
		}
	}

	return string(b)
}

func isVowel(r byte) bool {
	v := []byte("aiueoAIUEO")
	return slices.Contains(v, r)
}

func main() {
	fmt.Println("345-reverse-vowels-of-string")

	testCases := []struct {
		s        string
		expected string
	}{
		{s: "Hello", expected: "Holle"},
		{s: "leetcode", expected: "leotcede"},
		{s: "IceCreAm", expected: "AceCreIm"},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: s: %v\n", tc.s)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", reverseVowels(tc.s))
		fmt.Printf("--------\n\n")
	}
}
