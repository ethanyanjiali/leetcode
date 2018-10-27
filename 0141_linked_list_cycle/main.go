package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	// if fast.Next == nil and fast != slow, there can't be a cycle
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	n4 := ListNode{0, nil}
	n3 := ListNode{0, &n4}
	n2 := ListNode{0, &n3}
	n1 := ListNode{0, &n2}
	n4.Next = &n1
	fmt.Println(hasCycle(&n1))
	n4.Next = nil
	fmt.Println(hasCycle(&n1))
}
