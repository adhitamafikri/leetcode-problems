package main

import (
	"fmt"
)

// using binary search
func searchInsert(nums []int, target int) int {
	length := len(nums)
	low, high := 0, length-1

	if nums[length-1] < target {
		return length
	}

	for high >= low {
		diff := high - low
		if diff%2 != 0 {
			diff = diff/2 - 1
		}
		mid := (low + diff) / 2

		if nums[mid] == target {
			return mid
		}
		// move the space to left
		if nums[mid] > target {
			high = mid - 1
		}
		// move the space to right
		if nums[mid] < target {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println("35. Search Insert Position")

	testCases := []struct {
		input    []int
		target   int
		expected int
	}{
		{input: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{input: []int{1, 3, 5, 6}, target: 2, expected: 1},
		{input: []int{1, 3, 5, 6}, target: 7, expected: 4},
	}

	for index, tc := range testCases {
		fmt.Printf("Executing case #%d:\nInput: %v\ntarget: %d\nexpected: %d\nactual: %d\n", index, tc.input, tc.target, tc.expected, searchInsert(tc.input, tc.target))
	}

}
