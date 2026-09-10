package main

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。
示例 1：
输入：nums = [1,1,1], k = 2
输出：2
示例 2：
输入：nums = [1,2,3], k = 3
输出：2
提示：
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
解题思路:
首先: 连续子数组和的问题，首先考虑用前缀和来实现
前缀和数组s是一个长度为len(nums) + 1
s[0] = 0
s[i] = nums[0] + nums[1] + ... + nums[i-1]
再考虑两数之和问题:
s[j] - s[i]: 代表nums[i] 到nums[j-1]的元素和为k
*/
// 560_和为K的子数组: 前缀和
func subarraySum1(nums []int, k int) int {
	l := len(nums)
	ans := 0
	// 创建一个前缀和的哈希表
	preSumMap := make(map[int]int, l+1)
	preSumMap[0] = 1
	// 定义一个前缀和变量
	prefixSum := 0
	for _, num := range nums {
		prefixSum += num
		if preSumMap[prefixSum-k] > 0 {
			ans += preSumMap[prefixSum-k]
		}
		preSumMap[prefixSum]++
	}
	return ans
}

func subarraySum2(nums []int, k int) int {
	l := len(nums)
	// 创建一个结果变量
	ans := 0
	// 创建一个前缀和两数之和的哈希表
	prefixSumMap := make(map[int]int, l+1)
	// 前缀和遍历
	prefixSum := 0
	for _, num := range nums {
		prefixSumMap[prefixSum]++
		prefixSum += num
		ans += prefixSumMap[prefixSum-k]
	}
	return ans
}

func subarraySum3(nums []int, k int) int {
	l := len(nums)
	// 创建一个结果变量
	ans := 0
	// 创建一个前缀和的数组
	prefixSumList := make([]int, l+1)
	// 计算前缀和
	for idx, num := range nums {
		prefixSumList[idx+1] = prefixSumList[idx] + num
	}
	// 创建一个两数之和的哈希表
	m := make(map[int]int, l+1)
	// 计算两数之和
	for _, sj := range prefixSumList {
		ans += m[sj-k]
		m[sj]++
	}
	return ans
}
