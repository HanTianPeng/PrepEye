package main

/*
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符
（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
示例 1：
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
示例 2：
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
示例 3：
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。
提示：
1 <= text1.length, text2.length <= 1000
text1 和 text2 仅由小写英文字符组成。
解题思路:
对于718_最长重复子数组的区别在于这套题不要求连续
1. 确定dp含义及下标的定义: 整体全部最优(前缀型)
	dp[i][j]: 长度为[0,i-1]的字符串text1与长度为[0,j-1]的字符串text2的最长公共子序列
2. 确定递推公式
	if text1[i-1] == text2[j-1] dp[i][j] = dp[i-1][j-1] + 1
	if text[i-1] != text2[j-1] dp[i][j] = max(dp[i-1][j], dp[i][j-1])
3. 初始化
	dp[i][0]: text1[0, i-1]与空字符串的最长公共子序列为0
	dp[0][0]: text2[0, j-1]与空字符串的最长公共子序列为0
*/
// 1143_最长公共子序列: 动态规划二维数组
func longestCommonSubsequence(text1 string, text2 string) int {
	// 统计长度
	l1, l2 := len(text1), len(text2)
	// 创建动态规划二维数组
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	// 动态规划二维数组遍历
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

// 最长公共子序列: 动态规划一维数组
func longestCommonSubsequence2(text1 string, text2 string) int {
	// 统计长度
	l1, l2 := len(text1), len(text2)
	// 创建动态规划二维数组
	dp := make([]int, l2+1)
	// 动态规划二维数组遍历
	for i := 1; i <= l1; i++ {
		// dp[i-1][j-1]
		pre := 0
		for j := 1; j <= l2; j++ {
			cur := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = cur
		}
	}
	return dp[l2]
}
