package main

import (
	"fmt"
	// "math"
)

func search(nums []int, target int) int {
	length := len(nums)
	low, high, diff, mid := 0, length-1, 0, 0

	for high >= low {
		diff = high - low
		mid = low + (diff / 2)

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			high = mid - 1
		}

		if nums[mid] < target {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	type TestCase struct {
		Input    []int
		Target   int
		Expected int
	}

	testCases := []TestCase{
		{Input: []int{1, 3, 5, 6}, Target: 5, Expected: 2},
		{Input: []int{1, 3, 5, 6}, Target: 2, Expected: -1},
		{Input: []int{1, 3, 5, 6}, Target: 7, Expected: -1},
		{Input: []int{1, 3, 5, 6}, Target: 1, Expected: 0},
		{Input: []int{1, 3, 5, 6}, Target: 6, Expected: 3},
		{Input: []int{-1, 0, 3, 5, 9, 12}, Target: 9, Expected: 4},
		{Input: []int{-1, 0, 3, 5, 9, 12}, Target: 2, Expected: -1},
	}

	for i, tc := range testCases {
		actual := search(tc.Input, tc.Target)
		fmt.Printf("Case #%d | Input: %v | Target: %d | Expected: %d | Actual: %d\n", i+1, tc.Input, tc.Target, tc.Expected, actual)
	}
}
