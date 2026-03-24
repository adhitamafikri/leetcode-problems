package main

import "fmt"

func findTheDifference(s string, t string) byte {
	// Hint: Memoization
	// Memo data model:
	// type ocurrence struct { ocs: int, oct: int }
	// map[byte]occurrence

	// The idea is
	// 1. Get the length of both strings
	// 2. Take the longest length as reference
	// 3. Do for loop with the chosen length
	// 4. Record the occurrence of each characters to the memo
	// 5. Take any character which occurrence does not equal between the occurrence in s and t

	type occurrence struct {
		ocs int
		oct int
	}

	// memo := make(map[byte]int)
	memo := make(map[byte]*occurrence)

	l1, l2 := len(s), len(t)
	reference := 0

	if l1 < l2 {
		reference = l2
	} else {
		reference = l1
	}

	for i := 0; i < reference; i++ {
		// s
		if i < l1 {
			if _, ok := memo[s[i]]; !ok {
				memo[s[i]] = &occurrence{
					ocs: 1,
					oct: 0,
				}
			} else {
				memo[s[i]].ocs++
			}
		}

		// t
		if i < l2 {
			if _, ok := memo[t[i]]; !ok {
				memo[t[i]] = &occurrence{
					ocs: 0,
					oct: 1,
				}
			} else {
				memo[t[i]].oct++
			}
		}
	}

	// extract the character from map that have the occurrence exactly 1
	var result byte
	for key, value := range memo {
		if value.ocs != value.oct {
			result = key
			break
		}
	}

	return result
}

func main() {
	fmt.Println("389. Find the Difference")

	testCases := []struct {
		s        string
		t        string
		expected byte
	}{
		{s: "abcd", t: "abcde", expected: 'e'},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: s: %v, t: %v\n", tc.s, tc.t)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", findTheDifference(tc.s, tc.t))
		fmt.Printf("--------\n\n")
	}
}
