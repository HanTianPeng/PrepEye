package main

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
示例 1：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：
输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
提示：
1 <= prices.length <= 105
0 <= prices[i] <= 104
解题思路:
1. 定义dp[i]的函数和下标:
	dp[i][0]: 第i天持有股票
	dp[i][i]: 第i天不持有股票
2. 确认动态转换方程式
	dp[i][0] = max(dp[i-1][0], dp[i-1][1] - prices[i])
	dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])
3. 初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0
4. 遍历顺序及判断条件
	从i=1开始遍历
*/

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	// 边界判断
	if len(prices) == 0 {
		return 0
	}
	// 创建动态转换方程式二维数组变量
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}
	// 初始化dp数组
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		// 第i天持有股票
		dp[i][0] = max(dp[i-1][0], -prices[i])
		// 第i天不持有股票
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func maxProfit2(prices []int) int {
	// 边界判断
	if len(prices) == 0 {
		return 0
	}
	// 初始化dp数组
	dp0 := -prices[0]
	dp1 := 0
	for i := 1; i < len(prices); i++ {
		// 第i天持有股票
		dp0 = max(dp0, -prices[i])
		// 第i天不持有股票
		dp1 = max(dp1, dp0+prices[i])
	}
	return dp1
}
