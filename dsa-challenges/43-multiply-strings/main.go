package main

import (
	"math/big"
)

func multiply(num1 string, num2 string) string {
	n1 := new(big.Int)
	n2 := new(big.Int)
	res := new(big.Int)

	n1.SetString(num1, 10)
	n2.SetString(num2, 10)
	res.Mul(n1, n2)

	return res.String()
}
