package main

// 84_柱状图的最大的矩形: 待办事项 + 单调递增栈
func largestRectangleArea(heights []int) int {
	// 避免heights只有一个元素的时候无法计算以当前为柱状图,推算出左边和右边比这个元素小的所构成的宽度
	heights = append(heights, -1)
	// 创建一个单调递增栈
	stack := []int{-1}
	ans := 0
	for rightMin, height := range heights {
		// 只有右边碰到了比栈顶矮的元素,才证明以当前为中心向右扩展计算宽度结束,左边元素是递增,所以左边元素扩散到栈顶的下面个元素即可
		for len(stack) > 1 && height < heights[stack[len(stack)-1]] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leftMin := stack[len(stack)-1]
			ans = max(ans, heights[cur]*(rightMin-leftMin-1))
		}
		stack = append(stack, rightMin)
	}
	return ans
}
