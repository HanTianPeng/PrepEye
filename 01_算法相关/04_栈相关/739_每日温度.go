package main

// 739_每日温度: 单调递减栈 + 逆序遍历
func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	ans := make([]int, l)
	if l == 0 {
		return ans
	}
	// 创建一个单调递减栈
	stack := make([]int, 0)
	// 逆序遍历
	for i := l - 1; i >= 0; i-- {
		// 维持单调递减栈
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			// 栈的缩容
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

// 739_每日温度: 待处理栈 + 正序遍历
func dailyTemperatures1(temperatures []int) []int {
	l := len(temperatures)
	ans := make([]int, l)
	if l == 0 {
		return ans
	}
	// 创建一个待处理栈
	stack := make([]int, 0)
	for i, temperature := range temperatures {
		// 处理栈: 从栈末尾弹出一个元素,temperature大于这个元素,则证明能计算出结果
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			ans[j] = i - j
			stack = stack[:len(stack)-1]
		}
		// 默认全部进入到待处理栈
		stack = append(stack, i)
	}
	return ans
}
