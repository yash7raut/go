package main

import "fmt"

func main()  {
	
	nums := [] int{2,3,4}
	sum := 0
	for index,num := range nums{
		sum += num
		fmt.Println("index:",index)
	}

	fmt.Println("sum:",sum)
	//fmt.Println("index:",index)

	for i,	num := range nums{
		if num == 3{
			fmt.Println("index:",i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s -> %s\n",k,v)
	}

	for k := range kvs{
		fmt.Println("key:",k)
	}

	for a,b := range "go"{
		fmt.Println(a,b)
	}
}