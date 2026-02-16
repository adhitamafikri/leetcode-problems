// Naive way
function twoSumNaive(numbers: number[], target: number): number[] {
  const len = numbers.length
  let lp = 0
  let rp = lp + 1
  const result: number[] = []

  for (let i = 0; i < len; i++) {
      for (let j = i + 1; j < len; j++) {
          if (numbers[i] + numbers[j] === target) {
              return [i + 1, j + 1]
          }
      }
  }

  return result
};

// Optimal way
function twoSum(numbers: number[], target: number): number[] {
  const len = numbers.length
  let lp = 0
  let rp = len - 1
  const result: number[] = []

  // sum > target -> move the right ptr to left
  // sum < target -> move the left ptr to right

  let sum = 0
  while(true) {
      sum = numbers[lp] + numbers[rp]
      if (sum === target) {
          return [lp+1, rp+1]
      }

      if (sum > target) {
          rp--
      }

      if (sum < target) {
          lp++
      }

      if (lp >= rp) {
          break
      }
  }

  return result
};