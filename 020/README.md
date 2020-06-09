题目：[有效的括号](./https://leetcode-cn.com/problems/valid-parentheses)

这道题需要判断括号是否正确闭合，则右括号必须与左边最近的左括号匹配

可以模拟一个栈结构，后进的先出（最后进栈的左括号先与右括号匹配）

匹配成功出栈，失败返回 false