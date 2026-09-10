package main

/*
给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
每步 可以删除任意一个字符串中的一个字符。
示例 1：
输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
示例  2:
输入：word1 = "leetcode", word2 = "etco"
输出：4
提示：
1 <= word1.length, word2.length <= 500
word1 和 word2 只包含小写英文字母
解题思路:
1. 确定dp[i][j]定义及下标的含义
	dp[i][j]:以下标i-1为结尾的word1和以下标j-1为结尾的word2相等,所需要删除元素的最少个数dp[i][j]
2. 递推公式的推导
	当word1[i-1]=word2[j-1]: dp[i][j] = dp[i-1][j-1]
	当word1[i-1]!=word2[j-1]: dp[i][j]就是下面的取最小值
		1.删除word1: dp[i-1][j]+1
		2. 删除word2: dp[i][j-1]+1
		3. 删除word1,word2: dp[i-1][j-1]+2
3. 初始化
	dp[i][0]=i 删除i个元素
	dp[0][j]=j 删除j个元素
*/

// 583_两个字符串的删除操作: 动态规划二维数组
func minDistance(word1 string, word2 string) int {
	// dp[i][j]:以下标i-1为结尾的word1和以下标j-1为结尾的word2相等,所需要删除元素的最少个数dp[i][j]
	l1, l2 := len(word1), len(word2)
	// 创建动态规划二维数组并初始化
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	for j := 0; j < l2+1; j++ {
		dp[0][j] = j
	}
	// 动态规划二维数组遍历
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
			}
		}
	}
	return dp[l1][l2]
}

// 583_两个字符串的删除操作: 动态规划一维数组: 求出两个字符串的1143_最长公共子序列
func minDistance2(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([]int, l2+1)
	for i := 1; i < l1+1; i++ {
		pre := 0
		for j := 1; j < l2+1; j++ {
			cur := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = cur
		}
	}
	return l1 + l2 - dp[l2]*2
}
