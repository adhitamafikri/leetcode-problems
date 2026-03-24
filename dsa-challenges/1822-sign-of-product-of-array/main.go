package main

import (
	"fmt"
	"math/big"
)

func arraySign(nums []int) int {
	// Hint: finding product of an array -> simply multiply items within the array
	// e.g. [1, 2, 3, 4], the product is 1 * 2 * 3 * 4 = 24
	// Note that there is a possibility of having a big int result

	// result := 1
	result := big.NewInt(1)
	for _, num := range nums {
		result.Mul(result, big.NewInt(int64(num)))
	}

	return signFunc(result)
}

func signFunc(x *big.Int) int {
	return x.Sign()
}

func main() {
	fmt.Println("1822. Sign of Product of Array")

	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{-1, -2, -3, -4, 3, 2, 1}, expected: 1},
		{input: []int{1, 5, 0, 2, -3}, expected: 0},
		{input: []int{-1, 1, -1, 1, -1}, expected: -1},
		{input: []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}, expected: -1},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: %v\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", arraySign(tc.input))
		fmt.Printf("--------\n\n")
	}
}
