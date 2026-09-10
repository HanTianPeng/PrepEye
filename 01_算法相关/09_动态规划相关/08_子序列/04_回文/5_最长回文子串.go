package main

/*
给你一个字符串 s，找到 s 中最长的 回文 子串。
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：
输入：s = "cbbd"
输出："bb"
提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
解题思路:
	参考: 647_回文子串
*/
// 5_最长回文子串: 动态规划二维数组 + 647_回文子串
func longestPalindrome(s string) string {
	l := len(s)
	// 创建动态规划数组
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
	}
	// 创建结果变量
	maxLen, ans := 0, ""
	// 动态规划二维维度遍历
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
			if dp[i][j] && maxLen < j-i+1 {
				maxLen = j - i + 1
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
