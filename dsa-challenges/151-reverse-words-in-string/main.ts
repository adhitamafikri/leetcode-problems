function reverseWords(s: string): string {
    // Ideas:
    // - Remove extra spaces: trailing, leading, extra in between
    // - Reverse the word order (4 to 1, 3 to 2, 2 to 3, 1 to 4)

    // Steps:
    // 1. Trim the string
    // 2. Fragment the string by " ", normalize by trimming each items again after that
    // 3. Loop from the last item of the fragments, compose the result string

    const fragments = s.split(" ").map(frag => frag.trim()).filter(frag => frag !== "")
    const fragLength = fragments.length

    let result = ""
    for (let i = 0; i < fragLength; i++) {
        result = fragments[i] + result

        // check if the next item exists
        if (fragments[i + 1]) {
            result = " " + result
        }
    }

    return result
};

function main() {
  console.log("151-reverse-words-in-string");

  const testCases = [
    { input: "the sky is blue", expected: "" },
    { input: "  hello world  ", expected: "" },
    { input: "a good   example", expected: "" },
  ];

  let result = "";
  testCases.forEach((tc, index) => {
    result = reverseWords(tc.input);
    console.log(`Test Case #${index}:`);
    console.log(`Input: ${tc.input}`);
    console.log(`Expected: ${tc.expected}`);
    console.log(`Actual: ${result}`);
    console.log(`--------\n`);
  });
}

main();
