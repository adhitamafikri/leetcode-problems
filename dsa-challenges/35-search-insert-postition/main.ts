function searchInsert(nums: number[], target: number): number {
  const length = nums.length;
  let low = 0;
  let high = length - 1;

  if (nums[length - 1] <= target) {
    return length;
  }

  let diff = 0;
  let mid = 0;
  while (high >= low) {
    diff = high - low;
    if (diff > 2) {
      mid = low + Math.floor(diff / 2);
    } else {
      mid = low
    }

    console.log("iter", mid, low, high);

    if (nums[mid] === target) {
      return mid;
    }

    if (nums[mid] > target) {
      high = mid - 1;
    } else {
      low = mid + 1;
    }

    console.log("iter result: ", mid, low, high);
  }

  return 0;
}

function main() {
  console.log("35. Search Insert Position");

  const testCases = [
    { input: [1, 3, 5, 6], target: 5, expected: 2 },
    { input: [1, 3, 5, 6], target: 2, expected: 1 },
    { input: [1, 3, 5, 6], target: 7, expected: 4 },
    // { input: [1, 3, 5, 6], target: 0, expected: 0 },
    // { input: [3, 5, 7, 9, 10], target: 8, expected: 3 },
  ];

  testCases.forEach((tc, index) => {
    console.table({
      "Case #": index + 1,
      Input: tc.input.toString(),
      Target: tc.target,
      Expected: tc.expected,
      Actual: searchInsert(tc.input, tc.target),
    });
  });
}

main();
