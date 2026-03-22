package main

import "fmt"

func isAnagram(s string, t string) bool {
	// Hint: hash map, odd-even, struct
	// 1. Get the length of both strings
	// 2. Take the longest length as reference to be used in for loop
	// 3. Iterate the whole characters from both of strings and put the character and its occurrence into the map
	// 4. Iterate the map, if there is a character having ODD number of occurrence, then it is not valid anagram

	type occurrence struct {
		ocs int // ocurrence in string s
		oct int // ocurrence in string t
	}

	dict := make(map[string]*occurrence)
	sl, tl := len(s), len(t)
	ref := 0

	// edge case length is not the same, can return false immediately
	if sl != tl {
		return false
	}

	if sl < tl {
		ref = tl
	} else {
		ref = sl
	}

	for i := 0; i < ref; i++ {
		if i < sl {
			if _, ok := dict[string(s[i])]; !ok {
				dict[string(s[i])] = &occurrence{
					ocs: 1,
					oct: 0,
				}
			} else {
				dict[string(s[i])].ocs++
			}
		}

		if i < tl {
			if _, ok := dict[string(t[i])]; !ok {
				dict[string(t[i])] = &occurrence{
					ocs: 0,
					oct: 1,
				}
			} else {
				dict[string(t[i])].oct++
			}
		}
	}

	// check if it's valid anagram by iterating the map
	for _, value := range dict {
		if value.ocs != value.oct {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("28. Find the Index of the First Occurrence in a String")

	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "anagram", t: "nagaram", expected: true},
		{s: "kasur", t: "rusak", expected: true},
		{s: "basquiat", t: "hugoboss", expected: false},
		{s: "Fiasco", t: "Fias", expected: false},
		{s: "mexico", t: "oximec", expected: true},
	}

	result := false
	for index, tc := range testCases {
		result = isAnagram(tc.s, tc.t)
		fmt.Printf("Test Case #%d: %v, expected: %v\n", index, result, tc.expected)
	}
}
