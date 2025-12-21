function addBinary(a: string, b: string): string {
  const numA = BigInt(`0b${a}`);
  const numB = BigInt(`0b${b}`);
  const sum = numA + numB;

  return sum.toString(2);
}

function main() {
  const testCases = [
    { inputA: "11", inputB: "1", expected: "100" },
    { inputA: "1010", inputB: "1011", expected: "10101" },
  ];

  testCases.forEach((tc, index) => {
    console.log(`Case #${index}:`);
    console.log(`InputA: ${tc.inputA}`);
    console.log(`InputB: ${tc.inputB}`);
    console.log(`Expected: ${tc.expected}`);
    console.log(`Actual: ${addBinary(tc.inputA, tc.inputB)}`);
    console.log("--------");
  });
}

main();
