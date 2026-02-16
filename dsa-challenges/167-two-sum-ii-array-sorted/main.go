// Naive way
package main

func twoSum(numbers []int, target int) []int {
	length := len(numbers)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{}
}
