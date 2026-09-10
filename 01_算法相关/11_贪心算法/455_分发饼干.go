package main

import "sort"

// 分发饼干: 先排序 + 贪心算法 + 小饼干先分小胃口
func findContentChildren(g []int, s []int) int {
	// 先对两个数组进行排序
	sort.Ints(g) // nlogn
	sort.Ints(s) // mlogm

	// 定义一个结果变量
	ans := 0

	// 小饼干优先满足小胃口的小孩=> m
	gIndex := 0
	for i := 0; i < len(s); i++ {
		if gIndex < len(g) && s[i] >= g[gIndex] {
			ans++
			gIndex++
		}
	}
	return ans
}
