> 最近一直在 `LeetCode` 上刷题来熟悉 Go 的语法，遇见一道很绕头的题分享一下

题目要求：反转一个单链表

题目示例：
```
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
```
题目解法：
```Go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
```

> 答案解析

初始状态时：
![file](../statics/Reverse%20Linked%20List/01.png)

让我们迭代至 `head` 指向 `4` 时，`res = 5 -> nil`
我们要使 `5 -> nil` 变为 `5 -> 4`，则 `head.Next.Next = head`，即
![file](../statics/Reverse%20Linked%20List/02.png)

此时明显可以看出 `4和5` 形成了一个环，则 `head.Next = nil`，即
![file](../statics/Reverse%20Linked%20List/03.png)

此时可知当前返回后迭代至 `head` 指向 `3` 时 `res = 5 -> 4 -> nil`
同样的根据上述流程可得
![file](../statics/Reverse%20Linked%20List/04.png)

最终结果
![file](../statics/Reverse%20Linked%20List/05.png)