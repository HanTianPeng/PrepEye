package main

import "sort"

// 40_组合总和二: 去重 + 排序
func combinationSum2(candidates []int, target int) [][]int {
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
			// 裁剪
			if sum+candidates[i] > target {
				break
			}
			// 去重
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i + 1)
		}
	}
	// 排序方便后续进行裁剪
	sort.Ints(candidates)
	dfs(0)
	return ans
}
