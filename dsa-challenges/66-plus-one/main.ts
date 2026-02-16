function plusOne(digits: number[]): number[] {
  const joined = digits.join('')
  const converted = `${BigInt(joined) + BigInt(1)}`
  const result: number[] = []

  for (const ch of converted) {
      result.push(parseInt(ch, 10))
  }

  return result
};
