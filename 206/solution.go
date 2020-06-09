package main

import (
	"github.com/Han-MeiM/leetCode/structure/ListNode"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode.ListNode) *ListNode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func main() {
	root := ListNode.CreateTestData("[1,2,3,4,5]")
	res := reverseList(root)
	ListNode.Print(res)
}
