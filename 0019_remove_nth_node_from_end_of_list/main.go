package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	if head != nil {
		fmt.Printf("%d", head.Val)
		for head.Next != nil {
			head = head.Next
			fmt.Printf(" -> %d", head.Val)
		}
	}
	fmt.Println()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    first := dummy
    second := dummy
    for i := 0; i < n; i++ {
    	if second.Next == nil {
    		return nil
    	}
    	second = second.Next
    }
    for second.Next != nil {
    	second = second.Next
    	first = first.Next
    }
    first.Next = first.Next.Next
    return dummy.Next
}

func main() {
	a3 := ListNode{ 3, nil }
	a2 := ListNode{ 2, &a3 }
	a1 := ListNode{ 1, &a2 }
	printList(removeNthFromEnd(&a1, 3))
}