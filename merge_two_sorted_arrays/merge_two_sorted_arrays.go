package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {    
    dummy := &ListNode{}
    current := dummy
    for list1 != nil && list2 != nil{
        if list1.Val < list2.Val{
            current.Next = list1
            list1 = list1.Next
        }else{
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }
    if list1 != nil{
        current.Next = list1
    } else{
        current.Next = list2
    }
    return dummy.Next
}

func main() {
    // Create two sorted lists
    list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
    list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}

    // Call the mergeTwoLists function
    merged := mergeTwoLists(list1, list2)

    // Print the merged list
    current := merged
    for current != nil {
        fmt.Printf("%d -> ", current.Val)
        current = current.Next
    }
    fmt.Println("nil")
}

