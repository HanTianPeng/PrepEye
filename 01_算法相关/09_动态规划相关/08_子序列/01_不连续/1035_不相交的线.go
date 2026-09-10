package main

/*
在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：

	nums1[i] == nums2[j]

且绘制的直线不与任何其他连线（非水平线）相交。
请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
以这种方法绘制线条，并返回可以绘制的最大连线数。
示例 1：
输入：nums1 = [1,4,2], nums2 = [1,2,4]
输出：2
解释：可以画出两条不交叉的线，如上图所示。
但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。
示例 2：
输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
输出：3
示例 3：
输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
输出：2
提示：
1 <= nums1.length, nums2.length <= 500
1 <= nums1[i], nums2[j] <= 2000
解题思路:

	本质就是nums1和nums2的最长公共子序列的长度; 参考1143题
*/
// 1035_不相交的线: 动态规划二维数组 + 1143_最长公共子序列
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// 统计数组的长度
	l1, l2 := len(nums1), len(nums2)
	// 创建动态规划二维数组
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	// 动态规划二维数组初始化 dp[0][j]=0 dp[i][0]=0
	// 动态规划二维数组遍历
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

// 1035_不相交的线: 动态规划一维数组 + 1143_最长公共子序列
func maxUncrossedLines2(nums1 []int, nums2 []int) int {
	// 统计数组的长度
	l1, l2 := len(nums1), len(nums2)
	// 创建动态规划一维数组
	dp := make([]int, l2+1)
	// 初始化动态规划数组dp[0] = 0
	// 动态规划的一维数组遍历
	for i := 1; i < l1+1; i++ {
		pre := 0
		for j := 1; j < l2+1; j++ {
			cur := dp[j]
			if nums1[i-1] == nums2[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = cur
		}
	}
	return dp[l2]
}
