题目：[旋转链表](https://leetcode-cn.com/problems/rotate-list)

这道题思路是先找到链表长度 `len` ，然后再将 `tail` 与 `head` 相连形成一个环，再找到正确的地方将这个环断开即可，关键就是如何找到这个断点

这里采取从 `tail` 再往前走 `len - (k % len)` 步即为断点处

从上面的思路开始 `while` 循环获取链表长度 `len` ，并在循环到末尾处时将 `tail` 与 `head` 相连

因为这里是从 `head` 开始往前走，因此循环 `len - (k % len) - 1` 步即可

走完后将当前节点 `Next` 指向 `nil` 地址即可断开链表