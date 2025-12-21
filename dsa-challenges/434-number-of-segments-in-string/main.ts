function countSegments(s: string): number {
  const len = s.length;
  let segments = 0;

  if (len === 1) {
    if (s[0] !== " ") {
      return 1;
    }
    return 0;
  }

  let hasAddedSegment = false;
  for (let i = 0; i < len; i++) {
    if (s[i] === " ") {
      hasAddedSegment = false;
    } else {
      if (!hasAddedSegment) {
        segments++;
        hasAddedSegment = true;
      }
    }
  }

  return segments;
}

function main(): void {
  const testCases = [
    { input: "Hello, my name is John", expected: 5 },
    { input: "Hello", expected: 1 },
    { input: "                ", expected: 0 },
    { input: "a, b, c", expected: 3 },
  ];

  testCases.forEach((tc, index) => {
    console.log(`Case #${index}:`);
    console.log(`Input: ${tc.input}`);
    console.log(`Expected: ${tc.expected}`);
    console.log(`Actual: ${countSegments(tc.input)}`);
    console.log("--------");
  });
}

main();

export {};
