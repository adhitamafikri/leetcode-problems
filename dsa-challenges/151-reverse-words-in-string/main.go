package main

import "fmt"

func reverseWords(s string) string {
	// 1. Get the length of the og string
	// 2. Prepare empty result var for constructing the full reversed sentence
	// 3. Prepare empty temp var for constructing each words
	// 4. Iterate the whole characters in the og string
	// 5. For each character, append to temp if the current character is alphanum, skip if the current char is " " and only prepend to the result var if temp is populated, empty temp afterwards
	// 6. Repeat the iteration
	// 7. Return the result variable

	l := len(s)
	result, temp := "", ""

	for i := 0; i < l; i++ {
		if string(s[i]) == " " {
			if temp == "" {
				continue
			} else {
				if result == "" {
					result = temp
					temp = ""
				} else {
					result = temp + " " + result
					temp = ""
				}
			}
		} else {
			temp += string(s[i])
		}
	}

	// prepend to result for the last time if the temp is not empty
	if temp != "" {
		if result == "" {
			result = temp
			temp = ""
		} else {
			result = temp + " " + result
			temp = ""
		}
	}

	return result
}

func main() {
	fmt.Println("151. reverse words in string")

	testCases := []struct {
		input    string
		expected string
	}{
		{input: "the sky is blue", expected: "blue is sky the"},
		{input: "  hello world  ", expected: "world hello"},
		{input: "a good   example", expected: "example good a"},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: %v\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", reverseWords(tc.input))
		fmt.Printf("--------\n\n")
	}
}
