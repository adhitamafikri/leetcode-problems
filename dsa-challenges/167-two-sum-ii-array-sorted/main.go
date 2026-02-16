package main

// Naive way
func twoSumNaive(numbers []int, target int) []int {
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

// Optimal way
func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	lp := 0
	rp := length - 1

	sum := 0
	for {
		sum = numbers[lp] + numbers[rp]

		if sum == target {
			return []int{lp + 1, rp + 1}
		}

		if sum > target {
			rp--
		}

		if sum < target {
			lp++
		}

		if lp >= rp {
			break
		}
	}

	return []int{}
}
