package main

// 131_分隔回文串: 回文字符串 + 回溯
func partition(s string) [][]string {
	path, ans := make([]string, 0), make([][]string, 0)
	// 定义一个是否回文字符串的函数: 双指针
	var isPalindrome func(x string) bool
	isPalindrome = func(x string) bool {
		if len(x) == 0 {
			return true
		}
		left, right := 0, len(x)-1
		for left <= right {
			if x[left] != x[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	// 定义一个回溯递归函数
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i < len(s); i++ {
			//获取子串
			str := s[start : i+1]
			// 判断是不是回文串
			if !isPalindrome(str) {
				continue
			}
			path = append(path, str)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
