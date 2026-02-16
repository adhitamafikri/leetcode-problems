package main

func findDuplicate(nums []int) int {
	length := len(nums)
	lp, rp := 0, length-1
	cache := make(map[int]int)

	// best case
	if nums[lp] == nums[rp] {
		return nums[lp]
	}

	for {
		if length%2 == 0 {
			if lp >= rp {
				break
			}
		}

		if lp == rp {
			if _, ok := cache[lp]; ok {
				cache[nums[lp]]++
			} else {
				cache[nums[lp]] = 1
			}

			if cache[nums[lp]] > 1 {
				return nums[lp]
			}
		}

		if _, ok := cache[nums[lp]]; ok {
			cache[nums[lp]]++
		} else {
			cache[nums[lp]] = 1
		}

		if _, ok := cache[nums[rp]]; ok {
			cache[nums[rp]]++
		} else {
			cache[nums[rp]] = 1
		}

		// try to return the value
		if cache[nums[lp]] > 1 {
			return nums[lp]
		}

		if cache[nums[rp]] > 1 {
			return nums[rp]
		}

		// moving the pointers
		lp++
		rp--
	}

	return -1
}
