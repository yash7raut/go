package main

import (
	"fmt"
	"time"
)

func main() {

	i := 3
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

		switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's a Weekend")
		default:
			fmt.Println("It's a Weekday")
		}
	}
}
