package main

/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数。
输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
解题思路:
完全背包问题:
 1. 定义dp[i][j]及下标的含义: i种类型硬币,组合成硬币金额总和为j的方式有多少种
 2. 初始化:
    dp[0][j]: j % coins[0]如果余数为0，则dp[0][j] = 1
    dp[i][0]: 1
 3. 状态转换方程式
    dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
 4. 一维数组的遍历顺序: 组合,不讲究顺序; 所以一定要进行->先遍历商品，再遍历背包
*/
// 一维数组: 先遍历商品,再遍历背包
func changeV4(amount int, coins []int) int {
	// 创建动态规划一维数组:dp[j] amount为j的时候,组合有多少个
	dp := make([]int, amount+1)
	// 初始化
	dp[0] = 1
	// 动态规划遍历: 先遍历商品,再遍历背包
	for _, coin := range coins {
		// j := coin; j <= amount; j++ 不建议这种写法,因为coins数组没有明确说明是递增数列
		for j := 1; j <= amount; j++ {
			// 如果背包的重量小于商品的重量: dp[j]不能处理
			if j >= coin {
				dp[j] += dp[j-coin]
			}
		}
	}
	return dp[amount]
}

// 二维数组: 先遍历商品,再遍历背包
func change(amount int, coins []int) int {
	// 创建动态规划二维数组
	goodsLen := len(coins)
	dp := make([][]int, goodsLen)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		// 初始化当背包为0的时候的边界值
		dp[i][0] = 1
	}
	// 初始化当i=0的时候的边界值
	coin0 := coins[0]
	for j := 1; j <= amount; j++ {
		if j%coin0 == 0 {
			dp[0][j] = 1
		}
	}
	// 动态规划遍历: 先遍历商品再遍历背包
	for i := 1; i < goodsLen; i++ {
		for j := 1; j <= amount; j++ {
			// 如果背包的重量小于商品的重量
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}
		}
	}
	return dp[goodsLen-1][amount]
}

// 二维数组: 先遍历背包,再遍历商品
func changeV2(amount int, coins []int) int {
	// 创建动态规划二维数组
	goodsLen := len(coins)
	dp := make([][]int, goodsLen)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		// 初始化当背包重量为0时候的边界值
		dp[i][0] = 1
	}
	// 初始化当i=0的边界值
	for j := 1; j <= amount; j++ {
		// 如果能够整除，证明有一种方案
		if j%coins[0] == 0 {
			dp[0][j] = 1
		}
	}
	// 动态规划遍历: 先遍历背包,再遍历商品
	for j := 1; j <= amount; j++ {
		// 最后遍历商品
		for i := 1; i < goodsLen; i++ {
			// 如果背包的重量小于商品的重量
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}
		}
	}
	return dp[goodsLen-1][amount]
}
