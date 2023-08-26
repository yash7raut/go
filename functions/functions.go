package main

import "fmt"

func plus(a string , b string) string{
	return a + b
}

func plusPlus(a,b,c int) int{
	return a+b+c
}

func main(){

	res := plus("1", "2")
	fmt.Println("1+2 =",res)

	mes := plusPlus(1,2,3)
	fmt.Println("1+2+3 =",mes)

	fmt.Println(plus("1","2"))
}