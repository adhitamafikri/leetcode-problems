package main

import "fmt"

func strStr(haystack string, needle string) int {
	// Hint: two pointer, only spawn 2nd pointer when we found first character of the needle in the haystack
	// 1. Iterate the whole characters of haystack
	// 2. Check if the current character is the first letter of the needle
	// 3. If yes, set the second pointer to the i + needle_length
	// 4. Check if the slice of [i:i+needle_length] matches the needle
	// 5. If the result matches needle, immediately return the i
	// 6. Otherwise, continue looping from the ptr as the new i
	// 7. If no result until the end of the loop, it's guaranteed to return -1
	// Note: immediately return -1 if the haystack is shorter than the needle

	hl, nl := len(haystack), len(needle)
	ptr := 0

	// edge case: needle is longer than haystack
	if hl < nl {
		return -1
	}

	// normal cases
	for i := 0; i < hl; i++ {
		if string(haystack[i]) == string(needle[0]) {
			ptr = i + nl

			if ptr > hl {
				return -1
			}

			if haystack[i:ptr] == needle {
				return i
			}
		}
	}

	return -1
}

func main() {
	fmt.Println("28. Find the Index of the First Occurrence in a String")

	testCases := []struct {
		haystack string
		needle   string
		expected int
	}{
		{haystack: "sadbutsad", needle: "sad", expected: 0},
		{haystack: "leetcode", needle: "leeto", expected: -1},
		{haystack: "This is the art of basquiat", needle: "basq", expected: 19},
		{haystack: "Bartholomew Fiasco", needle: "Fias", expected: 12},
		{haystack: "Barbara Fiat Fias", needle: "Fias", expected: 13},
		{haystack: "mississippi", needle: "issipi", expected: -1},
	}

	result := -1
	for index, tc := range testCases {
		result = strStr(tc.haystack, tc.needle)
		fmt.Printf("Test Case #%d: %d, expected: %d\n", index, result, tc.expected)
	}
}
