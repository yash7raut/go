package main

import "fmt"

func main(){
	x := 5
	ptr := &x

	fmt.Println("Value of x = ",x)
	fmt.Println("Memory location of x = ",ptr)
	fmt.Println("Memory location of x = ",&x)
	fmt.Println("value of &x = ",*ptr)
}