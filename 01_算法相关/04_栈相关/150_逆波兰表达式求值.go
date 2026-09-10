package main

import "strconv"

// 逆波兰表达式求值: 栈实现
func evalRPN(tokens []string) int {
	// 边界判断
	if len(tokens) == 0 {
		return 0
	}
	// 创建一个栈的空间
	stack := make([]int, 0)
	// 遍历
	for _, token := range tokens {
		// 将token转换成int
		num, err := strconv.Atoi(token)
		if err != nil {
			// 从栈中弹出两个元素
			peek1, peek2 := stack[len(stack)-1], stack[len(stack)-2]
			// 栈缩容
			stack = stack[:len(stack)-2]
			// 进行数学计算
			if token == "+" {
				stack = append(stack, peek2+peek1)
			} else if token == "-" {
				stack = append(stack, peek2-peek1)
			} else if token == "*" {
				stack = append(stack, peek2*peek1)
			} else {
				stack = append(stack, peek2/peek1)
			}
		} else {
			stack = append(stack, num)
		}
	}
	return stack[0]
}
