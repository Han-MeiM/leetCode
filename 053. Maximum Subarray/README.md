题目：[最大子序和](https://leetcode-cn.com/problems/maximum-subarray)

第一种解法：普通的 O(n) 解法

先定义求和变量 `sum` 以及存储最大值的变量 `max`

循环数组 `nums` ，其中关键处为

```go
if sum < 0 {
    sum = 0
}
```

当 `sum < 0` 的时候，再与前面相加所得的和总是在减小的

当 `sum > 0` 的时候，再与前面相加所得的和总是在增加的

所以此处重新赋值 `sum = 0`

最后比较 `max` 与 `sum` ，将较大值赋值给 `max`

