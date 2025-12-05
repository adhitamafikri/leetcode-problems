package main

import "fmt"

func sumDifference(left []int, right []int) uint {
	leftSum := 0
	rightSum := 0

	for _, value := range left {
		leftSum += value
	}
	for _, value := range right {
		rightSum += value
	}

	return uint(leftSum - rightSum)
}

func countPartitions(nums []int) int {
	result := 0
	length := len(nums)

	for i := 0; i < length; i++ {
		// produce left and right slice
		if i < length-1 {
			left := nums[0 : i+1]
			right := nums[i+1 : length]
			temp := sumDifference(left, right)
			if temp%2 == 0 {
				result++
			}
		}
	}

	return result
}

/**
 * Ref: https://leetcode.com/problems/count-partitions-with-even-sum-difference/description/
 */
func main() {
	fmt.Println("Hello Worlds")

	// Expected: 4
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6}))
}
