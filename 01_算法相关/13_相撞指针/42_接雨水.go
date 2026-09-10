package main

/*
解题思路:
假设是[1, 0, 2, 5, 8, 7 2]
1-0才会构成一个坑位,所能接的雨水为 1 - 0 = 1
7-2不可以构成一个坑位,所以接不到雨水
*/

// 双指针:相向而行的前提条件是谁的高度越短,谁往对方方向移动;最终相遇点为最高位置
func trap(height []int) int {
	// 边界处理: 将最左边和最右边都添加一个0的元素,作为上一轮的双边的最大值
	lMax, rMax := 0, 0

	// 创建双指针,相向而行的前提条件是谁短就向对方而行；最终相遇的点一定是柱子的最高位置
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		// 需要更新lMax,rMax两个临时变量
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		// 判断左右两边谁大
		if lMax < rMax {
			// 计算左边能接的雨水数
			ans += lMax - height[l]
			l++
		} else {
			// 计算右边能接的雨水数
			ans += rMax - height[r]
			r--
		}
	}
	return ans
}

func trap1(height []int) int {
	// 创建一个结果变量
	ans := 0
	// 创建一个栈
	stack := make([]int, 0)
	// 遍历
	for i, h := range height {
		// 出栈
		if i > 0 && height[height[i]] <= h {
			w := i - height[i]
			ans += w * height[stack[len(stack)-1]]
		}
		// 入栈
		stack = append(stack, i)
	}
	return ans
}
