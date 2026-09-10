package main

/*
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
示例 1：
输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。
示例 2：
输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
输出：5
提示：
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 100
解题思路:
审题: 子数组其实就是连续子序列
1. 确定dp含义及下标的定义: 以i,j结尾的连续长度(结尾型)
	dp[i][j]: 以下标i-1结尾的A和以下标j-1结尾的B,最长重复子数组的长度为dp[i][j]
2. 确定递推公式
	if nums1[i] == nums2[j] dp[i][j] = dp[i-1][j-1] + 1
3. dp数组如何初始化
	dp[0][j] = 0
	dp[i][0] = 0
*/
// 718_最长重复子数组: 二维数组
func findLength(nums1 []int, nums2 []int) int {
	// 统计长度
	l1, l2 := len(nums1), len(nums2)
	// 创建动态规划二维数组dp[i][j]: 以下标i-1结尾的A和以下标j-1结尾的B,最长重复子数组的长度为dp[i][j]
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	// 定义结果变量
	ans := 0
	// 动态规划二维数组遍历
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

// 最长重复子数组: 动态规划一维数组
func findLength2(nums1 []int, nums2 []int) int {
	// 统计长度
	l1, l2 := len(nums1), len(nums2)
	// 创建动态规划一维数组
	dp := make([]int, l2+1)
	// 定义结果变量
	ans := 0
	// 动态规划一维数组遍历
	for i := 1; i <= l1; i++ {
		// 内层遍历-倒序遍历
		for j := l2; j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				// 不相等要初始化为0
				dp[j] = 0
			}
			if ans < dp[j] {
				ans = dp[j]
			}
		}
	}
	return ans
}
