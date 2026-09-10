package main

/*
给你一个由不同整数组成的数组nums和一个目标整数target。请你从 nums 中找出并返回总和为 target 的元素排列的个数。
输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
解题思路:
完全背包问题
讲究顺序: 先遍历背包,再遍历商品
dp[i]
*/
func combinationSum4(nums []int, target int) int {
	// 创建滚动数组
	dp := make([]int, target+1)
	// 初始化
	dp[0] = 1
	// 动态规划遍历: 先遍历背包,再遍历商品
	for j := 1; j <= target; j++ {
		for _, i := range nums {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}
	return dp[target]
}
