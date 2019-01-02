## 题名：[移除元素](https://leetcode-cn.com/problems/remove-element/)

这里采用双指针的方法

循环 `nums` 为快指针

`m` 为慢指针记录新的数组长度

当 `n` 不等于给定值时我们就复制 `n` 到 `nums[m]` 并递增 `m` 直到 `nums` 循环到尾部结束