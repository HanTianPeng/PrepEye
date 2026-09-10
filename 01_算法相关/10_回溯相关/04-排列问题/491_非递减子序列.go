package main

// 491_非递减子序列
func findSubsequences(nums []int) [][]int {
	path, ans := make([]int, 0), make([][]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) >= 2 {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}
		used := make(map[int]bool, len(nums))
		for i := start; i < len(nums); i++ {
			// 同一层去重
			if used[nums[i]] {
				continue
			}
			// 保证递增
			if len(path) > 0 && path[len(path)-1] > nums[i] {
				continue
			}
			// 当当前层的数字标记
			used[nums[i]] = true
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
