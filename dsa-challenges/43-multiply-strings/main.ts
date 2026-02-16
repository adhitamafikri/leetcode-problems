function multiply(num1: string, num2: string): string {
  const n1Conv = BigInt(num1)
  const n2Conv = BigInt(num2)
  const result = BigInt(n1Conv * n2Conv)

  return `${result}`
};
