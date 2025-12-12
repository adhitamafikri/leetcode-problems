package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	leftPtr := 0
	rightPtr := length - 1
	record := make(map[int]int) // caches the unique numbers that we have found
	result := 0                 // stores the number of unique

	for leftPtr < rightPtr {
		// iterate once O(n), but check with left and right pointer
		leftNum := nums[leftPtr]
		rightNum := nums[rightPtr]

		if _, ok := record[leftNum]; !ok {
			record[leftNum] = 1
			result++
		}

		if _, ok := record[rightNum]; !ok {
			record[rightNum] = 1
			result++
		}

		leftPtr++
		rightPtr--
	}

	return result
}

func main() {
	fmt.Println("Remove Duplicates from Sorted Array")
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 4, 5, 6}))
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 5, 10, 10}))
}
