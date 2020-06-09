## 题名：[求众数](https://leetcode-cn.com/problems/majority-element/)

首先创建一个 `map` 切片 `count` 用来装各数字出现的次数  

循环 `nums` 参数记录值出现的次数

`map` 中未定义的索引也是有值的，这里是 `make(map[int]int)` 则值为0

判断改值出现的次数是否大于 `nums` 长度的一半