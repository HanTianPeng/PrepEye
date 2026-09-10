package main

/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
示例 1：
输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：
输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。
提示：
1 <= s.length <= 1000
s 仅由小写英文字母组成
解题思路:
1. 确定dp数组（dp table）以及下标的含义
	dp[i][j]：字符串s在[i, j]范围内最长的回文子序列的长度为dp[i][j]。
2. 确定递推公式
	s[i] = s[j]: dp[i][j] = dp[i-1][j-1] + 1
	s[i] != s[j]: dp[i][j] = max(dp[i+1][j], dp[i][j-1])
	例如: a  b  a  a  b  c
		 i              j
    由于s[i] != s[j] 所以就有两种情况 i, j-1    i+1, j
3. 初始化
	首先要考虑当i 和j 相同的情况，从递推公式：dp[i][j] = dp[i + 1][j - 1] + 2; 可以看出 递推公式是计算不到 i 和j相同时候的情况。
	所以需要手动初始化一下，当i与j相同，那么dp[i][j]一定是等于1的，即：一个字符的回文子序列长度就是1。
	其他情况dp[i][j]初始为0就行，这样递推公式：dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]); 中dp[i][j]才不会被初始值覆盖。
*/

// 516_最长回本子序列
func longestPalindromeSubseq(s string) int {
	l := len(s)
	// dp[i][j]：字符串s在[i, j]范围内最长的回文子序列的长度为dp[i][j]。
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][l-1]
}
