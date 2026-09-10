package main

// 216_组合总和三: 不可重复
func combinationSum3(k int, n int) [][]int {
	// 定义变量
	path, sum, ans := make([]int, 0, k), 0, make([][]int, 0)
	// 定义回溯递归函数
	var dfs func(start int)
	dfs = func(start int) {
		// 递归结束条件
		if k == len(path) {
			// 路径不符合预期的结果直接剔除
			if sum != n {
				return
			}
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		// 树状遍历
		for i := start; i <= 9-(k-len(path))+1; i++ {
			// 剪枝: 1-和已经超过了n  2-数量不符合k个
			if sum+i > n {
				break
			}
			sum += i
			path = append(path, i)
			dfs(i + 1)
			// 恢复现场
			sum -= i
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
