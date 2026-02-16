function findDuplicate(nums: number[]): number {
  const length = nums.length
  let lp = 0    
  let rp = length - 1
  const cache: Record<string, number> = {}
  let result = 0

  while(true){
      if (length % 2 === 0) {
          if (lp >= rp) {
              break
          }
      }

      if (lp === rp) {
          if (!cache[nums[lp]]) {
              cache[nums[lp]] = 1
          } else {
              cache[nums[lp]]++
          }
          if (cache[nums[lp]] > 1) return nums[lp]
      }

      if (nums[lp] === nums[rp]) {
          return nums[rp]
      }

      // store the number and the occurrence to cache
      if (!cache[nums[lp]]) {
          cache[nums[lp]] = 1
      } else {
          cache[nums[lp]]++
      }

      if (!cache[nums[rp]]) {
          cache[nums[rp]] = 1
      } else {
          cache[nums[rp]]++
      }

      if (cache[nums[lp]] > 1) return nums[lp]
      if (cache[nums[rp]] > 1) return nums[rp]

      // move the pointers
      lp++
      rp--
  }

  return result
};
