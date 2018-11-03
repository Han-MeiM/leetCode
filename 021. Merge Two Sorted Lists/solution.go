package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var temp *ListNode
	if l1.Val < l2.Val {
		temp = l1
		temp.Next = mergeTwoLists(l1.Next, l2)
	} else {
		temp = l2
		temp.Next = mergeTwoLists(l1, l2.Next)
	}
	return temp
}

// 输出链表结果
func printList(l *ListNode) []int {
	val := l.Val
	temp := []int{val}
	if l.Next != nil {
		affter := printList(l.Next)
		temp = append(temp, affter...)
	}
	return temp
}

func main() {
	l1 := ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	res := mergeTwoLists(&l1, &l2)
	fmt.Println(printList(res))
}
