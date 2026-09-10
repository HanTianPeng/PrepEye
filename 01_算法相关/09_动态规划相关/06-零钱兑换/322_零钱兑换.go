package main

import "math"

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。
示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：
输入：coins = [2], amount = 3
输出：-1
示例 3：
输入：coins = [1], amount = 0
输出：0
解题思路:
 1. 确认dp数组及下标含义
    dp[j]: 凑足总金额为j所需钱币的最少个数为dp[j]
 2. 确定递推公式
    凑足金额为j-coin[i]所需金币的最少个数+1
    dp[j] = min(dp[j], dp[j-coin[i]]+1)
 3. 递推公式初始化
    dp[0] = 0
    dp[j] = math.MaxInt32 初始化最大值
 4. 确定遍历顺序
    求钱币的最少个数，有顺序和没有顺序都不会影响最终结果;所以本题遍历顺序无关;
    完全背包,内循环为正序
*/
// 零钱兑换: 先遍历商品,再遍历背包
func coinChange(coins []int, amount int) int {
	// 创建动态规划一维数组: 滚动数组
	dp := make([]int, amount+1)
	// 初始化: 因为取最小值,所以初始化为最大值
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	// 动态规划遍历-先遍历商品,再遍历背包
	for _, coin := range coins {
		for j := 1; j <= amount; j++ {
			// j >= coin也可以直接从j := coin开始遍历
			if j >= coin && dp[j-coin] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coin]+1)
			}
		}
	}
	// 没有找到能装满背包的组合
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 零钱兑换: 先遍历背包,再遍历商品
func coinChange2(coins []int, amount int) int {
	// 创建动态规划一维数组: 滚动数组
	dp := make([]int, amount+1)
	// 初始化
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	// 动态规划遍历-先遍历背包,再遍历商品
	for j := 1; j <= amount; j++ {
		for _, coin := range coins {
			if j >= coin && dp[j-coin] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
