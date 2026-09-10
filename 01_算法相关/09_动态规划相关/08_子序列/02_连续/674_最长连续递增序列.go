package main

/*
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
示例 1：
输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
示例 2：
输入：nums = [2,2,2,2,2]
输出：1
解释：最长连续递增序列是 [2], 长度为1。
提示：
1 <= nums.length <= 104
-109 <= nums[i] <= 109
解题思路:
 1. 确定dp定义及下标的定义: 以i结尾的最长连续子序列长度(结尾型)
    dp[i]: 以nums[i]为结尾的连续递增子序列长度
 2. 确定递推公式
    连续递增子序列讲究连续性,所以不需要j
    位置i的连续递增子序列等于以i-1为结尾的连续递增子序列长度+1
    if(nums[i] > nums[i-1]) dp[i] = dp[i-1] + 1
*/
// 最长连续递增序列: 动态规划一维数组
func findLengthOfLCIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans

}
