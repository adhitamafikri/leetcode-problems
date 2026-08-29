package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	// create map with -> key: array of alphabet occurrence (26 slots, lowercase only), value: array of strings
	// work on the naive way first with nested loop, where each loop will decide in which group the current word belongs to
	// notes:
	// ascii 97 to 122 (a-z)

	group := make(map[[26]int][]string)

	for _, word := range strs {
		var temp [26]int

		// scrutinize each characters from the word
		for _, r := range word {
			id := r - 97
			temp[id]++
		}

		// put the word into the group
		group[temp] = append(group[temp], word)
	}

	// form the result as nested array
	var result [][]string
	for _, val := range group {
		result = append(result, val)
	}

	return result
}

func main() {
	type testCase struct {
		input    []string
		expected [][]string
	}

	tc := []testCase{
		{input: []string{"act", "pots", "tops", "cat", "stop", "hat"}, expected: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}}},
		{input: []string{""}, expected: [][]string{{""}}},
		{input: []string{"a"}, expected: [][]string{{"a"}}},
	}

	for _, test := range tc {
		fmt.Println(test.input)
		fmt.Println(groupAnagrams(test.input))
	}
}
