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