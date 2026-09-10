package main

import (
	"sort"
)

/*
解题思路:
[[1,3],[2,6],[8,10],[15,18]]
合并后:
[[1,6],[8,10],[15,18]]
首先应该考虑的是以子数组的左端进行排序，后续直接合并即可
*/
// 合并区间: 排序 + 贪心算法
func merge(intervals [][]int) [][]int {
	// 定义变量
	res := make([][]int, 0)
	// 首先对二元数组进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 开始遍历
	for _, interval := range intervals {
		// 合并只有一种情况,res存在并且res的最后一个数组的1号索引的值大于等于interval的0号索引的值
		l := len(res)
		if l > 0 && res[l-1][1] >= interval[0] {
			res[l-1][1] = max(res[l-1][1], interval[1])
		} else {
			res = append(res, interval)
		}
	}
	return res
}
