package main

import "fmt"

// Naive way
func countOddsNaive(low int, high int) int {
	// Hint: Looping from the low, all the way to high
	// Adding 2 to any odd number will produce odd number

	var result []int

	for i := low; i <= high; i++ {
		if isOdd(i) {
			result = append(result, i)
		}
	}

	return len(result)
}

func isOdd(num int) bool {
	return num%2 != 0
}

// optimized way
func countOdds(low int, high int) int {
	// Hint: using math formula ((high - low)/2) + 1

	if isOdd(low) || isOdd(high) {
		return ((high - low) / 2) + 1
	}

	return ((high - low) / 2)
}

func main() {
	fmt.Println("1532. Count Odd Numbers in an Interval Range")

	testCases := []struct {
		low      int
		high     int
		expected int
	}{
		{low: 3, high: 7, expected: 3},
		{low: 15, high: 21, expected: 4},
		{low: 191, high: 217, expected: 14},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Low: %v\n", tc.low)
		fmt.Printf("High: %v\n", tc.high)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", countOdds(tc.low, tc.high))
		fmt.Printf("--------\n\n")
	}
}
