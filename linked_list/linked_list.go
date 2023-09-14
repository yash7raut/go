package main

import "fmt"

type Node struct{
	data int
	next *Node
}

func main(){
	//create node
	node1 := &Node{data: 1, next:nil}
	node2 := &Node{data: 2, next:nil}
	node3 := &Node{data: 3, next:nil}

	//linking nodes to form a list
	node1.next = node2
	node2.next = node3

	//traversing the linked list
	current := node1
	for current != nil{
		fmt.Println(current.data)
		current = current.next
	}

}