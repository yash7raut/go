package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	closeToOpen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if _, exists := closeToOpen[c]; exists {
			if len(stack) > 0 && stack[len(stack)-1] == closeToOpen[c] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()")) // Output: true
}
