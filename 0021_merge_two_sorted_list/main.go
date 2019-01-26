package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := dummy
	for l1 != nil || l2 != nil {
		// to optimize, we could break the loop here when one of the node is nil
		if l1 == nil || l2 != nil && l2.Val <= l1.Val {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		} else if l2 == nil || l1 != nil && l1.Val <= l2.Val {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		}
	}
	return dummy.Next
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

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	printList(mergeTwoLists(l1, l2))
}
