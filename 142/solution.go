package main

import (
	"fmt"
	"github.com/Han-MeiM/leetCode/structure/ListNode"
)

//func detectCycle(head *ListNode.ListNode) *ListNode.ListNode {
//	record := make(map[int]int)
//	for head != nil {
//		if _, ok := record[head.Val]; ok {
//			return head
//		}
//		record[head.Val] = 1
//		head = head.Next
//	}
//	return nil
//}

func detectCycle(head *ListNode.ListNode) *ListNode.ListNode {
	if head == nil {
		return nil
	}
	fast, slow, entry := head, head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for slow != entry {
				slow = slow.Next
				entry = entry.Next
			}
			return entry
		}
	}
	return nil
}

func main() {
	list := ListNode.CreateTestData("[3,2,0,4]")
	list.Next.Next.Next.Next = list.Next
	res := detectCycle(list)
	fmt.Println(res)
}
