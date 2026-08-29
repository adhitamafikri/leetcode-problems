package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Solution using map - naive solution
func compress(chars []byte) int {
	// Flow:
	// 1. Iterate all characters within the array
	// 2. For each character, record them into the map -> char x occurrence
	// 3. Increment the occurrence of a character if the next iteration is still the same char
	// 4. From the record, begin building the result string
	// 5. Remove the string length
	// Hint:
	// 1. The input is guaranteed to be sorted
	// 2. Utilize map[byte]int -> contains character and its occurrence
	// 3. Count of occurrence does not need to be appendded to the array if it's only 1

	record := make(map[byte]int)
	var resultBuilder strings.Builder

	for i := 0; i < len(chars); i++ {
		_, ok := record[chars[i]]
		if ok {
			record[chars[i]]++
		} else {
			record[chars[i]] = 1
		}
	}

	for char, occurrences := range record {
		resultBuilder.WriteByte(char)
		if occurrences > 1 {
			resultBuilder.WriteString(strconv.Itoa(occurrences))
		}
	}

	result := resultBuilder.String()
	return len(result)
}

func main() {
	fmt.Println("443-string-compression")

	testCases := []struct {
		input    []byte
		expected int
	}{
		{input: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, expected: 6},
		{input: []byte{'a'}, expected: 1},
		{input: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, expected: 4},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: %v\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", compress(tc.input))
		fmt.Printf("--------\n\n")
	}
}
