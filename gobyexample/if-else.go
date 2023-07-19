package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is an even number")
	} else {
		fmt.Println("7 is a odd number")
	}

	if num := 12; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
