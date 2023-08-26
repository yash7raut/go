package main

import "fmt"

func palindrome(x int) bool {
	if x < 0 {
		// negetive numbers cant be palindrome numbers
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return original == reversed
}

func main() {
	fmt.Println(palindrome(919))
}
