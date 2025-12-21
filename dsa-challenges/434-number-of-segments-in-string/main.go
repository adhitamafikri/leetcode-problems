package main

import "fmt"

func countSegmentConvertToChar(s string) int {
	segment := 0
	hasAddedSegment := false

	for _, item := range s {
		if string(item) != " " {
			if !hasAddedSegment {
				segment++
				hasAddedSegment = true
			}
		} else {
			hasAddedSegment = false
		}
	}

	return segment
}

func countSegments(s string) int {
	segment := 0
	hasAddedSegment := false

	for _, item := range s {
		if item != 32 {
			if !hasAddedSegment {
				segment++
				hasAddedSegment = true
			}
		} else {
			hasAddedSegment = false
		}
	}

	return segment
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "Hello, my name is John", expected: 5},
		{input: "Hello", expected: 1},
		{input: "                ", expected: 0},
		{input: "a, b, c", expected: 3},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: %s\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", countSegments(tc.input))
		fmt.Println("--------\n")
	}
}
