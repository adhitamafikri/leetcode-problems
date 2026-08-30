package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Solution using map - naive solution
func compressV1(chars []byte) int {
	// Flow:
	// 1. Iterate all characters within the array
	// 2. For each character, record them into the map -> char x occurrence
	// 3. Increment the occurrence of a character if the next iteration is still the same char
	// 4. From the record, begin building the result string
	// 5. Remove the string length
	// Hint:
	// 1. The input is guaranteed to be sorted
	// 2. Utilize map[byte]int -> contains character and its occurrence
	// 3. Count of occurrence does not need to be appendded to the array if it's only 1

	record := make(map[byte]int)
	var resultBuilder strings.Builder

	for i := 0; i < len(chars); i++ {
		_, ok := record[chars[i]]
		if ok {
			record[chars[i]]++
		} else {
			record[chars[i]] = 1
		}
	}

	for char, occurrences := range record {
		resultBuilder.WriteByte(char)
		if occurrences > 1 {
			resultBuilder.WriteString(strconv.Itoa(occurrences))
		}
	}

	result := resultBuilder.String()
	return len(result)
}

// Solution using map and in-place modification
func compress(chars []byte) int {
	// Solution breakdown:
	// 1. Iterate each character in the bytes array
	// 2. For each iteration, simply record the occurrence of the character that, as long as the current character is still the same with the initial character being pointed at
	// 3. Once the different character is found, replace the consecutive characters with the number of occurrences of that character, then point to the new character afterwards
	// 4. Repeat the process until the end

	// Hints
	// 1. Use two pointers tech, slow and fast pointer.
	// 2. Slow pointer is used to keep track of the position of the character, fast pointer is for iterating and counting the ocurrences

	slow := 0
	fast := slow + 1

	// use this pointer to decide in which index should we write the item to replace
	// why +1 from slow? Because the digits will be written immediately next to the current character pointed at by slow pointer
	writePointer := slow + 1

	// simple int variable to keep track of the occurrences of the current character that is being iterated
	occurrences := 1

	// edge case: only 1 item in byte array
	if len(chars) == 1 {
		return 1
	}

	// normal case
	for fast < len(chars) {
		// if the next character is still the same, just increase the occurrence
		if chars[slow] == chars[fast] {
			occurrences++
		} else {
			// if the next character is not the same, decide if we need to do in-place modification
			if occurrences > 1 {
				// convert the occurrences to byte array, to decide how many spaces we need to take for overwriting the digit into the array
				digits := []byte(strconv.Itoa(occurrences))

				// write the digits and advance the writePointer
				for d := 0; d < len(digits); d++ {
					chars[writePointer] = digits[d]
					writePointer++
				}

				// write the new character to iterate, next to the digit we've just written into the array
				chars[writePointer] = chars[fast]
				// move the write pointer next to the last index we've just overwritten
				writePointer++
			} else {
				chars[writePointer] = chars[fast]
				writePointer++
			}

			// move the slow to the next char, as the char on the slow pointer does not have any consecutive same characters
			slow = fast
			// set occurrences to 1, marking the number of occurrences of the new character on slow pointer
			occurrences = 1
		}

		// keep the fast pointer moving after each iteration
		fast++
	}

	// handle the final occurrence writing as the loop ends when the iteration ends
	// if the next character is not the same, decide if we need to do in-place modification
	if occurrences > 1 {
		// convert the occurrences to byte array, to decide how many spaces we need to take for overwriting the digit into the array
		digits := []byte(strconv.Itoa(occurrences))

		// write the digits and advance the writePointer
		for d := 0; d < len(digits); d++ {
			chars[writePointer] = digits[d]
			writePointer++
		}
	}

	// inspect the array
	fmt.Println("bytes after mod: ", chars)
	fmt.Println("actual word after mod: ", string(chars))
	fmt.Println("actual word after mod | compressed: ", string(chars[:writePointer]))

	// // cheat: just return the last position of the write pointer
	// return writePointer

	// actually return the compressed array by slicing until the final position of writePointer
	return len(chars[:writePointer])
}

func main() {
	fmt.Println("443-string-compression")

	testCases := []struct {
		input    []byte
		expected int
	}{
		{input: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, expected: 6},
		{input: []byte{'a'}, expected: 1},
		{input: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, expected: 4},
		{input: []byte{'a', 'a'}, expected: 2},
		{input: []byte{'a', 'a', 'a', 'b', 'b', 'c'}, expected: 5},
		{input: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c'}, expected: 6},
		{input: []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c'}, expected: 7},
		{input: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'a', 'b', 'c'}, expected: 12},
		{input: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'g', 'g', 'g', 'a', 'b', 'c'}, expected: 11},
	}

	for index, tc := range testCases {
		fmt.Printf("Case #%d:\n", index)
		fmt.Printf("Input: %v\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Actual: %v\n", compress(tc.input))
		fmt.Printf("--------\n\n")
	}
}
