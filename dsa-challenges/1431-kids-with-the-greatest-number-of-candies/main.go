package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// 1. Loop through all items in the candies array
	// 2. Get the highest item in the array
	// 3. Loop through the array again to produce a new array that contains the boolean

	highest := 0
	for _, c := range candies {
		if highest < c {
			highest = c
		}
	}

	result := make([]bool, len(candies))
	for idx, c := range candies {
		if c+extraCandies >= highest {
			result[idx] = true
		} else {
			result[idx] = false
		}
	}

	return result
}

func main() {
	fmt.Println("1431-kids-with-the-greatest-number-of-candies")

	testCases := []struct {
		s        []int
		t        int
		expected []bool
	}{
		{s: []int{2, 3, 5, 1, 3}, t: 3, expected: []bool{true, true, true, false, true}},
		{s: []int{4, 2, 1, 1, 2}, t: 1, expected: []bool{true, false, false, false, false}},
		{s: []int{12, 1, 12}, t: 10, expected: []bool{true, false, true}},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: s: %v, t: %v\n", tc.s, tc.t)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", kidsWithCandies(tc.s, tc.t))
		fmt.Printf("--------\n\n")
	}
}
