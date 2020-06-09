题目：[相交链表](https://leetcode-cn.com/problems/intersection-of-two-linked-lists)

这道题我们先假设有这两个链表

```
ListA = a1 -> a2 -> a3 -> c1 -> c2 -> c3

ListB = b1 -> b2 -> b3 -> b4 -> c1 -> c2 -> c3
```

将两个链表互相追加在后面可得

```
ListA = a1 -> a2 -> a3 -> c1 -> c2 -> c3 -> b1 -> b2 -> b3 -> b4 -> c1 -> c2 -> c3

ListB = b1 -> b2 -> b3 -> b4 -> c1 -> c2 -> c3 -> a1 -> a2 -> a3 -> c1 -> c2 -> c3
```

可以发现他们的尾部都是 

```
c1 -> c2 -> c3
```

既是我们所求的相交处