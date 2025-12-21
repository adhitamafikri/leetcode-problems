function search(nums: number[], target: number): number {
  const len = nums.length;
  let low = 0;
  let high = len - 1;
  let diff = 0;
  let mid = 0;

  for (let i = 0; i < len; i++) {
    diff = low - high;
    mid = low + Math.floor((high - low) / 2);

    if (nums[mid] === target) {
      return mid;
    }

    if (nums[mid] < target) {
      low = mid + 1;
    }

    if (nums[mid] > target) {
      high = mid - 1;
    }
  }

  return -1;
}

function main() {
  const testCases = [
    { input: [1, 3, 5, 6], target: 5, expected: 2 },
    { input: [1, 3, 5, 6], target: 2, expected: -1 },
    { input: [1, 3, 5, 6], target: 7, expected: -1 },
    { input: [1, 3, 5, 6], target: 1, expected: 0 },
    { input: [1, 3, 5, 6], target: 6, expected: 3 },
    { input: [-1, 0, 3, 5, 9, 12], target: 9, expected: 4 },
    { input: [-1, 0, 3, 5, 9, 12], target: 2, expected: -1 },
  ];

  testCases.forEach((tc, index) => {
    console.table({
      "Case #": index + 1,
      Input: tc.input.toString(),
      Target: tc.target,
      Expected: tc.expected,
      Actual: search(tc.input, tc.target),
    });
  });
}

main();

export {};
