## 题名：[缺失数字](https://leetcode-cn.com/problems/missing-number)

首先 `nums` 序列是从 `0 .. n` 的连续数字组成的

当 `nums` 不缺失时 `nums` 所有索引也是从 `0 .. n` 

则 `索引之和 == 数字之和`

因此我们要获取缺失的数字就可以用 `索引之和 - 现有数字之和`

而数字缺失的情况下，索引也会缺失一位最末尾的值也就是 `len(nums)`

则 `索引之和 = 现有索引之和 + len(nums)`

所以 `缺失数字 = 现有索引之和 + len(nums) - 现有数字之和`

通过代码可知 `i - v` 为 `索引 - 数字`

```go
sum := len(nums)
for i, v := range nums {
	sum = sum + i - v
}
```

