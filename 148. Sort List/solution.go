package main

import "leetCode/structure/ListNode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode.ListNode) *ListNode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head.Next
	// 使用快慢指针寻找中点
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	l := head
	r := p.Next
	p.Next = nil
	return merge(sortList(l), sortList(r))
}

func merge(l, r *ListNode.ListNode) *ListNode.ListNode {
	res := new(ListNode.ListNode) //初始化一个新链表
	tmp := res
	for l != nil && r != nil {
		if l.Val < r.Val {
			tmp.Next = l
			l = l.Next
			tmp = tmp.Next
		} else {
			tmp.Next = r
			r = r.Next
			tmp = tmp.Next
		}
	}
	if l != nil {
		tmp.Next = l
	}
	if r != nil {
		tmp.Next = r
	}
	return res.Next
}

func main() {
	res := ListNode.CreateTestData("[4,2,1,3]")
	res = sortList(res)
	ListNode.Print(res)
}
