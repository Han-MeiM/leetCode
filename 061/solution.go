package main

import (
	"github.com/Han-MeiM/leetCode/structure/ListNode"
)

func rotateRight(head *ListNode.ListNode, k int) *ListNode.ListNode {
	if head == nil {
		return head
	}
	temp, len := head, 1
	for temp.Next != nil {
		temp = temp.Next
		len = len + 1
	}
	temp.Next = head
	k = k % len
	for i := 0; i < len-k-1; i++ {
		head = head.Next
	}
	temp = head
	head = head.Next
	temp.Next = nil
	return head
}

func main() {
	head := ListNode.CreateTestData("[1,2,3,4,5,6]")
	res := rotateRight(head, 10)
	ListNode.Print(res)
}
