package main

import "fmt"

func moveZeroes(nums []int) []int {
	// Hint: in-place operations, two pointer, swapping values
	// 1. Get the length of the input array
	// 2. Define ptr1 = 0 and ptr2 = 1, ptr1 should move first, followed by ptr2
	// 3. If a pointer finds a 0, it should stop. The other ptr should move until it finds a non-zero item
	// 4. Swap the value, then move the pointer to the next item
	// 5. Repeat the steps, until all the zeroes are placed on the tail

	l := len(nums)
	p1 := 0
	p2 := p1 + 1

	for i := 0; i < l; i++ {
		if nums[p1] == 0 {
			for {
				if p2 != l && nums[p2] != 0 {
					tmpz := nums[p1]
					nums[p1] = nums[p2]
					nums[p2] = tmpz
					break
				}
				p2++

				if p2 >= l {
					break
				}
			}
		}
		p1++
		p2 = p1 + 1
	}

	return nums
}

func main() {
	fmt.Printf("283. Move Zeroes\n\n\n")

	testCases := []struct {
		s        []int
		expected []int
	}{
		{s: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
		{s: []int{45192, 0, -659, -52359, -99225, -75991, 0, -15155, 27382, 59818, 0, -30645, -17025, 81209, 887, 64648}, expected: []int{45192, -659, -52359, -99225, -75991, -15155, 27382, 59818, -30645, -17025, 81209, 887, 64648, 0, 0, 0}},
	}

	for index, tc := range testCases {
		snapshot := make([]int, len(tc.s))
		copy(snapshot[:], tc.s[:])
		moveZeroes(tc.s)
		fmt.Printf("Test Case #%d: %v\nGot: %v\nexpected: %v\n\n\n", index, snapshot, tc.s, tc.expected)
	}
}
