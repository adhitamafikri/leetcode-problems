package main

import "fmt"

func productExceptSelf(nums []int) []int {
	// Hint:
	// Product of an array means, multiplication result of all items in the array

	// Caveats:
	// Each item[i] in the output represents the product that is produced without the nums[i] being accounted for

	result := make([]int, len(nums))

	for i, _ := range nums {
		product := 1
		for j, nInner := range nums {
			if i != j {
				product *= nInner
			}
		}
		result[i] = product
	}

	return result
}

func main() {
	fmt.Println("238-product-of-array-except-self")

	testCases := []struct {
		s        []int
		expected []int
	}{
		{s: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{s: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: s: %v\n", tc.s)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", productExceptSelf(tc.s))
		fmt.Printf("--------\n\n")
	}
}
