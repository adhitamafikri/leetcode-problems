package main

import "fmt"

func removeElement(nums []int, val int) int {
	// The point is removing the `val` from the array
	// Should do in-place array modification
	// Use two-pointer tech (slow and fast)
	// Slow as anchor for the current number which is eq to val
	// Fast as anchor for the candidate value for swap
	// Return the last position of slow pointer OR the length of the slice of the og array to last position of slow pointer

	sp := 0
	fp := sp
	arrLen := len(nums)
	needToRemove := false

	// edge case
	if arrLen == 1 {
		if nums[0] == val {
			return 0
		}

		return 1
	}

Outer:
	for fp < arrLen {
		// check if the element pointed at by fast pointer is the value to be removed
		if nums[fp] == val {
			needToRemove = true

			// advance the fast pointer to search for the value for overriding the item pointed by the slow pointer
			// do this set of things until the fast pointer finds a value that is not the one to be removed
			for needToRemove {
				// advance fast pointer
				fp++

				// escape hatch, terminate the loop if after fp increment it matches the length of the array
				if fp == arrLen {
					break Outer
				}

				// check if current pointed value is not the `val`
				if nums[fp] != val {
					// swap if the condition is met, then advance the slow pointer by 1, set the fast pointer to slow pointer
					temp := nums[sp]
					nums[sp] = nums[fp]
					nums[fp] = temp

					sp++
					fp = sp

					needToRemove = false
				} else {
					// terminate the loop if fp touches the last item but the value on fast pointer is still eq to `val`
					if fp == arrLen-1 {
						break Outer
					}
				}
			}
		} else {
			// if the number is not the one to remove just advance the pointers
			sp++
			fp++
		}
	}

	// // returning just the last position of slow pointer
	// fmt.Println("After modification:")
	// fmt.Println("nums: ", nums)
	// fmt.Println("last position of sp: ", sp)

	// return sp

	// returning the actual sliced array
	fmt.Println("After modification:")
	fmt.Println("nums: ", nums[:sp])
	fmt.Println("last position of sp: ", sp)

	return len(nums[:sp])
}

func main() {

	type testCase struct {
		nums        []int
		valToRemove int
		expected    int
	}
	tc := []testCase{
		{nums: []int{3, 2, 2, 3}, valToRemove: 3, expected: 2},
		{nums: []int{1, 2, 2, 3, 0, 4, 2}, valToRemove: 2, expected: 2},
		{nums: []int{4, 5}, valToRemove: 4, expected: 1},
	}

	for i, c := range tc {
		fmt.Printf("Case #%d:\n", i)
		fmt.Printf("Input: %v\n", c.nums)
		fmt.Printf("Expected: %v\n", c.expected)
		fmt.Printf("Actual: %v\n", removeElement(c.nums, c.valToRemove))
		fmt.Printf("--------\n\n")
	}
}
