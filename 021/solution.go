package main

import (
	"github.com/Han-MeiM/leetCode/structure/ListNode"
)

func mergeTwoLists(l1 *ListNode.ListNode, l2 *ListNode.ListNode) *ListNode.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var temp *ListNode.ListNode
	if l1.Val < l2.Val {
		temp = l1
		temp.Next = mergeTwoLists(l1.Next, l2)
	} else {
		temp = l2
		temp.Next = mergeTwoLists(l1, l2.Next)
	}
	return temp
}

func main() {
	l1 := ListNode.CreateTestData("[1,2,4]")
	l2 := ListNode.CreateTestData("[1,3,4]")
	res := mergeTwoLists(l1, l2)
	ListNode.Print(res)
}
