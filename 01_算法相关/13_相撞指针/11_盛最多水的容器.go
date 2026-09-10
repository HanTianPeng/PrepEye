package main

// 盛最多水的容器: 双向指针
func maxArea(height []int) int {
	// 创建双指针,左指针代表左边的高度,右指针代表右边的高度;谁最小;代表容器的高度
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		// 两个指针的差值代表容器的宽度
		area := (r - l) * (min(height[l], height[r]))
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
