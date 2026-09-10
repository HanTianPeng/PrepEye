package main

/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true
示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false
提示：
0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。
解题思路:
  参考1143_最长公共子序列
1. 确定dp[i][j]的定义及下标的含义
	dp[i][j]: 以下标i-1为结尾的字符串s和以下标j-1为结尾的字符串t,相等的子序列长度为dp[i][j]
2. 递推公式推导
	if s[i-1]==t[j-1]: dp[i][j] = dp[i-1][j-1] + 1
	if s[i-1]!=t[j-1]: dp[i][j] = dp[i][j-1] => 等同于t需要删除元素
3. 初始化
	dp[i][0]: 0; 子序列长度只能是0
	dp[0][j]: 0; 子序列长度只能是0
*/
// 392_判断子序列: 动态规划二维遍历 + 最长公共子序列
func isSubsequence(s string, t string) bool {
	l1, l2 := len(s), len(t)
	// 创建动态规划二维数组
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	// 动态规划二维数组遍历
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 因为是判断s是否为t的子序列,所以这意思代表的是删除t的元素
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[l1][l2] == l1
}

// 392_判断子序列: 动态规划一维遍历 + 最长公共子序列
func isSubsequence2(s string, t string) bool {
	// 统计长度
	l1, l2 := len(s), len(t)
	// 创建动态规划一维数组
	dp := make([]int, l2+1)
	// 动态规划遍历
	for i := 1; i < l1+1; i++ {
		pre := 0
		for j := 1; j < l2+1; j++ {
			cur := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = dp[j-1]
			}
			pre = cur
		}
	}
	return dp[l2] == l1
}

// 392_判断子序列: 动态规划二维数组 + bool判断 + 类似回文子串
func isSubsequence3(s string, t string) bool {
	l1, l2 := len(s), len(t)
	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}
	// 动态规划初始化
	for j := 0; j < l2+1; j++ {
		dp[0][j] = true
	}
	// 动态规划遍历
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[l1][l2]
}
