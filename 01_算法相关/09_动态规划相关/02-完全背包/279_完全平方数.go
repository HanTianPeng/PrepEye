package main

import "math"

/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
示例 1：
输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
示例 2：
输入：n = 13
输出：2
解释：13 = 4 + 9
提示：
1 <= n <= 104
解题思路:
完全背包问题
遍历顺序: 不讲究
*/
func numSquares(n int) int {
	// 创建动态规划一维数组:滚动数组
	dp := make([]int, n+1)
	// 初始化
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	// 动态规划遍历: 先遍历背包,再遍历商品
	for j := 1; j <= n; j++ {
		for i := 1; i <= j; i++ {
			if j >= i*i && dp[j-i*i] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}

// 完全背包: 先遍历商品再遍历背包
func numSquaresV2(n int) int {
	// 创建动态规划一维数组:滚动数组
	dp := make([]int, n+1)
	// 初始化
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	// 动态规划遍历: 先遍历背包,再遍历商品
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			if dp[j-i*i] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}
