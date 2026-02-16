package main

import (
	"math/big"
	"strconv"
)

func plusOne(digits []int) []int {
	// naive way
	// convert to string
	joined := ""
	for _, dig := range digits {
		joined += strconv.Itoa(dig)
	}

	// convert to Big Int using math/big and add by 1
	n := new(big.Int)
	n.SetString(joined, 10)
	n.Add(n, big.NewInt(1))

	var result []int
	for _, dig := range n.String() {
		result = append(result, int(dig)-'0')
	}

	return result
}
