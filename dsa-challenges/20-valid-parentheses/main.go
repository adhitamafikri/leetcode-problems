package main

import "fmt"

func checkBracketType(s string) string {
	if s == "(" || s == "{" || s == "[" {
		return "opening"
	}

	if s == ")" || s == "}" || s == "]" {
		return "closing"
	}

	return "invalid"
}

func isValid(s string) bool {
	var stack []string

	result := false
	for _, item := range s {
		ch := fmt.Sprintf("%c", item)
		if checkBracketType(ch) == "opening" {
			stack = append(stack, ch)
			result = false
		} else {
			length := len(stack)
			if length == 0 {
				return false
			}
			if ch == ")" {
				if stack[length-1] == "(" {
					stack = stack[:length-1]
					result = true
				} else {
					return false
				}
			}
			if ch == "}" {
				if stack[length-1] == "{" {
					stack = stack[:length-1]
					result = true
				} else {
					return false
				}
			}
			if ch == "]" {
				if stack[length-1] == "[" {
					stack = stack[:length-1]
					result = true
				} else {
					return false
				}
			}
		}
	}

	// return false if the stack is not fully emptied
	if len(stack) != 0 {
		return false
	}

	return result
}

func main() {
	fmt.Println("20. Valid Parentheses\n----------------")
	fmt.Println(isValid("()"))
	fmt.Println(isValid("[]{}()"))
	fmt.Println(isValid("[)"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("){"))
	fmt.Println(isValid("([]){"))
}
