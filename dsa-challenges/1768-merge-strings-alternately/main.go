package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	// Hint: Two pointers
	// Take the longest string length as reference
	// Using 1 'for' loop, iterate characters on both string simultaneously
	// Build the result value with each characters from both words

	len1, len2 := len(word1), len(word2)
	reference := 0

	if len1 < len2 {
		reference = len2
	} else {
		reference = len1
	}

	p1, p2 := 0, 0
	result := ""
	for i := 0; i < reference; i++ {
		if p1 < len1 {
			result += string(word1[p1])
			p1++
		}

		if p2 < len2 {
			result += string(word2[p2])
			p2++
		}
	}

	return result
}

func main() {
	testCases := []struct {
		id       int
		word1    string
		word2    string
		expected string
	}{
		{id: 1, word1: "abc", word2: "def", expected: "adbecf"},
		{id: 2, word1: "per", word2: "xxx", expected: "pxexrx"},
		{id: 3, word1: "tour", word2: "battlepass", expected: "tboautrtlepass"},
	}

	for _, tc := range testCases {
		fmt.Printf("Case #%d. %s %s\nGot %s, expected %s\n", tc.id, tc.word1, tc.word2, mergeAlternately(tc.word1, tc.word2), tc.expected)
	}
}
