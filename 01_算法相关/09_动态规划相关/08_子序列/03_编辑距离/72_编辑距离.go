package main

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：
输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
提示：
0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
解题思路:
1.dp[i][j]的定义及下标含义
	表示以下标i-1为结尾的字符串word1，和以下标j-1为结尾的字符串word2，最近编辑距离为dp[i][j]
2. 递推公式推导
  word1[i-1]与word2[j-1]相等: dp[i-1][j-1]
  word1[i-1]与word2[j-1]不相等: 取下面三种的最小值
	1. 删除word1一个元素: dp[i-1][j]
    2. 增加: dp[i][j-1] => 等同于删除word2
    3. 替换: dp[i-1][j-1]
*/
// 72_编辑距离: 动态规划二维数组 + 初始化dp[0][j],dp[i][0]
func minDistanceV1(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	// 创建动态规划二维数组
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	// 初始化
	for j := 1; j < l2+1; j++ {
		dp[0][j] = j
	}
	// 动态规划遍历
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[l1][l2]
}
