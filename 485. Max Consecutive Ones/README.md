## 思路

题名：[最大连续1的个数](https://leetcode-cn.com/problems/max-consecutive-ones/)

预设两个长度变量 `maxL` 和 `tempL` 

分别表示最大连续1的长度和当前连续 1 的长度

循环数组 `nums` 并判断当前值是否等于 1

如果等于 1 则 `tempL + 1` 并在之后判断 `tempL` 是否大于连续 1 的最大长度 `maxL` 

如果大于了就直接替换掉 `maxL`

若不等于 1 则将当前长度置为 0 