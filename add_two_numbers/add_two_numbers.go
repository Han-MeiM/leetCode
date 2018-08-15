package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *  求两个单链表的和
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{0, nil}
	p := l1
	q := l2
	curr := &dummyHead
	carry := 0
	for (p != nil || q != nil) {
		var x, y int

		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum := carry + x + y

		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}

		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return dummyHead.Next
}

/**
 * LeetCode 优质解答
 */
func addTwoNumbersLeetCode(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	p := res
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{carry % 10, nil}
		p = p.Next
		carry = carry / 10
	}
	if carry != 0 {
		p.Next = &ListNode{carry, nil}
	}
	return res.Next
}