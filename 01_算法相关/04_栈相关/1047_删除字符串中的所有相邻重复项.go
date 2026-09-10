package main

// 删除字符串中的所有相邻重复项
func removeDuplicates(s string) string {
	// 创建一个栈
	stack := make([]rune, 0)
	// 字符串遍历
	for _, r := range s {
		if len(stack) == 0 || stack[len(stack)-1] != r {
			stack = append(stack, r)
		} else {
			// 入栈
			stack = stack[:len(stack)-1]
		}
		
	}
	return string(stack)
}
