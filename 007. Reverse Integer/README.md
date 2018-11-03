题目：[反转整数](https://leetcode-cn.com/problems/reverse-integer)

本题以 `-123` 作为示例

颠倒他类似于从最末位开始取他的数字最后拼接成新的数

```golang
// 取最末尾值的方法等于对10取余
last = x % 10

// 将倒数第二个数字推到末尾值
x = x / 10

// 提升res的位数，将最新值作为个位插入其中
res = res*10 + last
```