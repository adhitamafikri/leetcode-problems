package main

import (
	"fmt"
)

// func isPalindromeNaive(x int) bool {
// 	strNum := strconv.Itoa(x)
// 	length := len(strNum)
// 	reversed := ""

// 	for i := 0; i < length; i++ {
// 		reversed += fmt.Sprintf("%c", strNum[length-1-i])
// 	}
// 	return strNum == reversed
// }

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	temp := x

	for temp != 0 {
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}

	return reverse == x
}

func main() {
	fmt.Println("Is Palindrome?: ", isPalindrome(353))
	fmt.Println("Is Palindrome?: ", isPalindrome(-1234))
	fmt.Println("Is Palindrome?: ", isPalindrome(5423))
	fmt.Println("Is Palindrome?: ", isPalindrome(97455479))
}
