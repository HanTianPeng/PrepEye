package main

import "sort"

// 90_子集二: 重复 + 树层去重 + 排序
func subsetsWithDup(nums []int) [][]int {
	path, ans := make([]int, 0), make([][]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		// start: 用来记录纵向递归的次数
		if start > len(nums) {
			return
		}
		// 得到子集
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		// 回溯遍历
		for i := start; i < len(nums); i++ {
			// 树层去重
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	// 首先进行排序
	sort.Ints(nums)
	dfs(0)
	return ans
}
