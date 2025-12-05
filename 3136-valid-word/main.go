package main

import "fmt"

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	hasVowel := false
	hasConsonant := false

	// play based on ascii
	// numbers (48-57)
	numAsciiMap := map[string]int{
		"0": 48,
		"1": 49,
		"2": 50,
		"3": 51,
		"4": 52,
		"5": 53,
		"6": 54,
		"7": 55,
		"8": 56,
		"9": 57,
	}
	// alphabet - uppercase (65-90)
	upperAsciiMap := map[string]int{
		"A": 65,
		"B": 66,
		"C": 67,
		"D": 68,
		"E": 69,
		"F": 70,
		"G": 71,
		"H": 72,
		"I": 73,
		"J": 74,
		"K": 75,
		"L": 76,
		"M": 77,
		"N": 78,
		"O": 79,
		"P": 80,
		"Q": 81,
		"R": 82,
		"S": 83,
		"T": 84,
		"U": 85,
		"V": 86,
		"W": 87,
		"X": 88,
		"Y": 89,
		"Z": 90,
	}
	// alphabet - lowercase (97-122)
	lowerAsciiMap := map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
		"d": 100,
		"e": 101,
		"f": 102,
		"g": 103,
		"h": 104,
		"i": 105,
		"j": 106,
		"k": 107,
		"l": 108,
		"m": 109,
		"n": 110,
		"o": 111,
		"p": 112,
		"q": 113,
		"r": 114,
		"s": 115,
		"t": 116,
		"u": 117,
		"v": 118,
		"w": 119,
		"x": 120,
		"y": 121,
		"z": 122,
	}

	vowelMap := map[string]int{
		"a": 97,
		"e": 101,
		"i": 105,
		"o": 111,
		"u": 117,
		"A": 65,
		"E": 69,
		"I": 73,
		"O": 79,
		"U": 85,
	}

	for _, item := range word {
		ch := string(item)
		_, okNum := numAsciiMap[ch]
		_, okUpper := upperAsciiMap[ch]
		_, okLower := lowerAsciiMap[ch]
		_, okVowel := vowelMap[ch]

		if !(okNum || okUpper || okLower) {
			return false
		}

		if okUpper || okLower {
			if okVowel {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		}
	}

	return hasConsonant && hasVowel
}

// Ref: https://leetcode.com/problems/valid-word
func main() {
	fmt.Println("is Word Valid: ", isValid("234Adas"))
	fmt.Println("is Word Valid: ", isValid("b3"))
	fmt.Println("is Word Valid: ", isValid("a3$e"))
}
