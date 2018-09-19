## 思路

题目：[删除链表中的节点](https://leetcode-cn.com/problems/delete-node-in-a-linked-list)

这道题描述有点不清楚

给你的参数是被删除的节点，而不是整个链表

所以我们只知道当前的节点和节点之后的内容

并不知道节点之前的内容

所以我们只需要将当前节点替换为下一个节点即可

```javascript
node.val = node.next.val

node.next = node.next.next
```