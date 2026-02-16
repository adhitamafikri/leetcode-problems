// Naive way
function twoSum(numbers: number[], target: number): number[] {
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
