package main

/*
给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。
测试用例保证结果在 32 位有符号整数范围内。
示例 1：
输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
示例 2：
输入：s = "babgbag", t = "bag"
输出：5
解释：
如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
babgbag
babgbag
babgbag
babgbag
babgbag
提示：
1 <= s.length, t.length <= 1000
s 和 t 由英文字母组成
解题思路:
1. 确定dp[i][j]的定义及下标的含义
  dp[i][j]是以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]
2. 递推公式的确定
  s[i-1]与t[j-1]相等时:
	用s[i-1]与t[j-1]进行匹配,那么就是dp[i-1][j-1]
    不用s[i-1]与t[j-1]进行匹配,那么就是dp[i-1][j]
  s[i-1]与t[j-1]不相等时:
    dp[i-1][j]
3. 初始化
	dp[i][0]=0 删除s的所有字符串等同于空字符串t，所以出现的次数为1
*/
// 115_不同的子序列: 动态规划二维数组
func numDistinct(s string, t string) int {
	// 统计长度
	l1, l2 := len(s), len(t)
	// 创建动态规划一维数组
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
		dp[i][0] = 1
	}
	// 动态规划遍历
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s[i-1] == t[j-1] {
				// dp[i-1][j-1]: 使用s[i-1]和t[j-1]匹配；dp[i-1][j]: 不适用s[i-1]和t[j-1]匹配
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				// 删除s字符串中的元素
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[l1][l2]
}

// 115_不同的子序列: 动态规划一维数组
func numDistinct2(s string, t string) int {
	l1, l2 := len(s), len(t)
	dp := make([]int, l2+1)
	dp[0] = 1
	for i := 1; i <= l1; i++ {
		// 一定要从后往前遍历
		for j := l2; j > 0; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[l2]
}
