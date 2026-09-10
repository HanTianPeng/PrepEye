package main

// 有效括号: 栈
func isValid(s string) bool {
	// 创建一个反向的map常量
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 创建一个栈
	stack := make([]rune, 0)
	// 字符串的遍历
	for _, r := range s {
		// 先将正向的入栈
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		peek := stack[len(stack)-1]
		// 判断是否还是拿去反向
		if peek != m[r] {
			return false
		}
		// 栈的缩容
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
