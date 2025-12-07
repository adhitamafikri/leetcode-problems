package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strNum := strconv.Itoa(x)
	length := len(strNum)
	reversed := ""

	for i := 0; i < length; i++ {
		reversed += fmt.Sprintf("%c", strNum[length-1-i])
	}

	fmt.Println(strNum, reversed)

	return strNum == reversed
}

func main() {
	fmt.Println("Is Palindrome?: ", isPalindrome(353))
	fmt.Println("Is Palindrome?: ", isPalindrome(-1234))
	fmt.Println("Is Palindrome?: ", isPalindrome(5423))
	fmt.Println("Is Palindrome?: ", isPalindrome(97455479))
}
