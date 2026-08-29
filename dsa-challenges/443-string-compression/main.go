package main

import "fmt"

func compress(chars []byte) int {
	// Hint: Utilize map[byte]int -> contains character and its occurrence
}

func main() {
	fmt.Println("443-string-compression")

	testCases := []struct {
		input    []byte
		expected int
	}{
		{input: []byte{'a','a','b','b','c','c','c'}, expected: 6},
		{input: []byte{'a'}, expected: 1},
		{input: []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}, expected: 4},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: %v\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", compress(tc.input))
		fmt.Printf("--------\n\n")
	}
}
