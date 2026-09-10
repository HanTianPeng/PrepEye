package main

import "sort"

// 39_组合总和: 可重复
func combinationSum(candidates []int, target int) [][]int {
	path, sum, ans := make([]int, 0), 0, make([][]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			// 前面已经进行排序后,这里直接可以剪枝
			if sum+candidates[i] > target {
				break
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i)
			// 恢复现场
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	// 首先进行排序
	sort.Ints(candidates)
	dfs(0)
	return ans
}
