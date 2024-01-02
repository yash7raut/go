package main

import (
	"fmt"
	"time"
)

func myFunction(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main(){
	go myFunction()

	time.Sleep(time.Second * 3)
	fmt.Println("Main fuction Exiting")
}