package main

func subsets(nums []int) [][]int {
	path, ans := make([]int, 0), make([][]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if start > len(nums) {
			return
		}
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		// 回溯for循环
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			// 恢复现场
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
