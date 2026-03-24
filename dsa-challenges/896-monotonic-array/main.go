package main

import "fmt"

func isMonotonic(nums []int) bool {
	// 1. Loop the whole nums array
	// 2. Check the first comparison (lt, gt)
	// 3. When the next comparison result between current and next item does not match the first comparison, return false immediately
	// 4. Else, finish the loop and return true

	comp := ""
	for idx, _ := range nums {
		if idx == 1 {
			if nums[idx] < nums[idx-1] {
				comp = "lt"
			} else if nums[idx] > nums[idx-1] {
				comp = "gt"
			}
		}

		if idx > 1 {
			if nums[idx] < nums[idx-1] && comp == "gt" {
				return false
			}

			if nums[idx] > nums[idx-1] && comp == "lt" {
				return false
			}

			if comp == "" {
				if nums[idx] < nums[idx-1] {
					comp = "lt"
				} else if nums[idx] > nums[idx-1] {
					comp = "gt"
				}
			}
		}
	}

	return true
}

func main() {
	fmt.Println("896. Monotonic Array")

	testCases := []struct {
		input    []int
		expected bool
	}{
		{input: []int{1, 2, 2, 3}, expected: true},
		{input: []int{6, 5, 4, 4}, expected: true},
		{input: []int{1, 3, 2}, expected: false},
		{input: []int{2, 2, 2, 1, 4, 5}, expected: false},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: %v\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", isMonotonic(tc.input))
		fmt.Printf("--------\n\n")
	}
}
