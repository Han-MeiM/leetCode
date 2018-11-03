## 思路

题名：[单调数列](https://leetcode-cn.com/problems/monotonic-array)

首先定义 `previous := 0` 来记录上一个数字是递增还是递减

从下标1也就是第二个数字开始循环数组 `A`

判断当前数字比上个数字大还是小（如果相等则直接跳过）

如果大则为递增且 `curren = 1`

如果小则为递减且 `curren = -1`

最后判断 `previous` 是否与 `curren` 相等

若 `previous` 为0则为初始状态并将 `curren` 赋值给它

