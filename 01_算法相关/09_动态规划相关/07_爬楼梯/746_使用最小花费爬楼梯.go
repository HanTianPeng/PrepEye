package main

/*
解题思路:
动态规划五步骤:
1. 确定dp数组的含义以及下标的含义
2. 确定递推公式
3. dp数组如何初始化
4. 确定遍历顺序
5. 举例推举dp数组
dp[i] = min(dp[i - 2] + cost[i - 2], dp[i - 1] + cost[i - 1])
dp[i]: 跑到i个台阶需要的最小花费
*/
func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}
	// 创建动态规划数组
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[len(cost)]
}
