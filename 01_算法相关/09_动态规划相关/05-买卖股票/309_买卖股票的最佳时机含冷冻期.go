package main

/*
给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
示例 1:
输入: prices = [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
示例 2:
输入: prices = [1]
输出: 0
提示：
1 <= prices.length <= 5000
0 <= prices[i] <= 1000
解题思路:

1. 确认动态转换方程式
	dp[i][0]: 第i天持有股票
	dp[i][1]: 第i天保持卖出股票
	dp[i][2]: 第i天卖出股票
	dp[i][3]: 第i天冷冻期
	1. 第i天达到买入股票状态: 前一天持有股票   前一天冷冻期今天买入   前一天保持卖出股票状态
	dp[i][0] = max(dp[i-1][0], dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])
	2. 第i天保持卖出股票状态: 前一天保持卖出股票状态  前一天是冷冻期
	dp[i][1] = max(dp[i-1][1], dp[i-1][3])
	3. 第i天达到今天卖出股票状态: 前一天一定是持有股票状态
	dp[i][2] = (dp[i-1][0]+prices[i])
	4. 第i天达到冷冻期状态: 前一天一定是卖出股票状态
	dp[i][3] = (dp[i-1][2])
*/

// 买卖股票的最佳时期含冷冻期
func maxProfitL1(prices []int) int {
	// 边界处理
	if len(prices) == 0 {
		return 0
	}
	// 创建动态规划二维数组
	dp := make([][]int, len(prices))
	for i, _ := range prices {
		if i == 0 {
			dp[0] = []int{-prices[0], 0, 0, 0}
		} else {
			dp[i] = make([]int, 4)
		}
	}
	// 动态规划遍历
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2], dp[len(prices)-1][3])
}
