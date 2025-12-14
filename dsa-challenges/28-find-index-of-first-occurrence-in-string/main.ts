function strStr(haystack: string, needle: string): number {
  return haystack.indexOf(needle)
}

function main() {
  console.log("28. Find the Index of the First Occurrence in a String");

  const testCases = [
    { haystack: "sadbutsad", neeedle: "sad", expected: 0 },
    { haystack: "leetcode", neeedle: "leeto", expected: -1 },
  ];

  let result = -1;
  testCases.forEach((tc, index) => {
    result = strStr(tc.haystack, tc.neeedle);
    console.log(`Test Case #${index}: ${result}`);
  });
}

main();
