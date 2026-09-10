package main

func generateParenthesis(n int) []string {
	path, ans := "", make([]string, 0)
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if len(path) == n*2 {
			ans = append(ans, path)
		}
		// 左边优先添加
		if left > 0 {
			left--
			path += "("
			dfs(left, right)
			left++
			path = path[:len(path)-1]
		}
		// 只有左括号的数量小于右括号的数量才能添加右括号
		if left < right {
			right--
			path += ")"
			dfs(left, right)
			right++
			path = path[:len(path)-1]
		}
	}
	dfs(n, n)
	return ans
}
