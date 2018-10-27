package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{val % 10, nil}
		carry = val / 10
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{carry, nil}
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
	a3 := ListNode{3, nil}
	a2 := ListNode{4, &a3}
	a1 := ListNode{2, &a2}
	b3 := ListNode{4, nil}
	b2 := ListNode{6, &b3}
	b1 := ListNode{5, &b2}
	res := addTwoNumbers(&a1, &b1)
	printList(res)
}
