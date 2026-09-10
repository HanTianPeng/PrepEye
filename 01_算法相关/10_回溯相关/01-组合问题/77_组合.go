package main

// 77_组合: 不可重复
func combine(n int, k int) [][]int {
	path, ans := make([]int, 0, k), make([][]int, 0)
	// 定义回溯递归函数
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(i + 1)
			// 恢复现场
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}

func combine2(n int, k int) [][]int {
	path, ans := make([]int, 0, k), make([][]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i <= n; i++ {
			// for循环的起始位置之后的元素个数已经不足我们需要的元素个数,直接进行剪枝操作
			if n-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			dfs(i + 1)
			// 恢复现场
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
