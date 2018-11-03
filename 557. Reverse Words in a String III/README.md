## 思路

题目：[反转字符串中的单词 III](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii)

首先获取字符串长度 `l`

从后往前开始循环并依次将当前字符串填入数组 `word` 中

空格 `ASCII` 码为 **32**

判断当前字符串是否等于空格

若等于则表示一个单词的结束

通过 `string` 函数将 `word` 数组组合为一个字符串

并将其拼接在结果变量 `res` 之前即可