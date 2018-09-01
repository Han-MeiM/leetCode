## 思路

题名：无重复字符的最长子串

首先从左开始循环作为子串的开头一直到末尾

```go
for index := range s
```



并同时循环以该开头以右的剩余字符串 `s[index+1:]` 

```go
for r_index, right := range s[index+1:]
```



则定义当前的子串为 `s[index:r_index + index +1]`



在循环中通过循环开头索引 `index` 和当前右侧字符串索引 `r_index + index + 1` 的值来判断当前字符串是否在子串中出现

```go
func in_str(s string, one int32) bool {
	result := false
	for _, v := range s {
		if one == v {
			result = true
			break
		}
	}
	return result
}
```



若出现过则结束当前循环，并以下一个字符串作为开头继续循环，否则当前子串长度加1 `sum += 1`



___



另一种解法 `imooc_programme`

`lastOccurred` 记录字符串最后出现的索引位置

`start` 记录子串开头索引位置

循环字符串并判断当前字符串的索引位置最后一次出现的地方 `lastI`

如果 `lastI` 大于子串开头位置 `start` 则表示当前字符串已经在子串中出现过

将 `start` 设置为 `lastI + 1` 

```go
// 当该字符串最后出现的位置在 start 和 i 之间时
if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
	start = lastI + 1
}
```

最后判断当前子串长度是否大于之前的最长子串长度并更新当前字符串最后出现的位置

```go
if i - start + 1 > maxLength {
	maxLength = i - start + 1
}

lastOccurred[ch] = i
```

