package main

// 17_电话号码的字母组合
func letterCombinations(digits string) []string {
	// 将数字与字母进行映射
	letterList := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	s, ans := "", make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(s) == len(digits) {
			ans = append(ans, s)
			return
		}
		// 获取当前的数字并转换为数字
		digit := digits[start] - '0'
		// 获取数字对应的字母
		letters := letterList[digit]
		for i := 0; i < len(letters); i++ {
			s += string(letters[i])
			dfs(start + 1)
			// 恢复现场
			s = s[:len(s)-1]
		}
	}
	dfs(0)
	return ans
}
