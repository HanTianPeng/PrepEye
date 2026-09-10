package main

import "sort"

// 47_全排列二: nums为可重复元素
func permuteUnique(nums []int) [][]int {
	l := len(nums)
	path, used, ans := make([]int, 0, l), make([]bool, l), make([][]int, 0)
	var dfs func()
	dfs = func() {
		if len(path) == l {
			temp := make([]int, l)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		// 每一层都是从其实位置开始提取元素
		for i := 0; i < len(nums); i++ {
			// 提取重复元素
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			// 已经取过的元素不再重复提取
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			// 恢复现场
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	// 去重首先考虑排序
	sort.Ints(nums)
	dfs()
	return ans
}
