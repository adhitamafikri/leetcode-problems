package main

import "fmt"

func strStr(haystack string, needle string) int {
	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	// edge case
	if lenNeedle > lenHaystack {
		return -1
	}

	result := -1 // the first index of the substring's occurrence
	bucket := 0  // stores how may characters we've matched with the needle
	for index, hs := range haystack {
		if string(hs) == string(needle[0]) {
			result = index
			// iterate through the needle
			for needleIndex, nd := range needle {
				if index+needleIndex >= lenHaystack {
					return -1
				}
				if string(nd) == string(haystack[index+needleIndex]) {
					bucket++
				}
			}

			if bucket == lenNeedle {
				return result
			} else {
				result = -1
				bucket = 0
			}
		}
	}

	return result
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
