package main

/*
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
示例 1：
输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：
输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
提示：
1 <= nums.length <= 200
1 <= nums[i] <= 100
解题思路:
典型的01背包问题:
背包的重量就是: sum/2
商品: nums
背包能装的产品最大价值和如果等于sum/2;则证明可以拆分成两个子集
*/
func canPartition(nums []int) bool {
	// 计算数组里面元素的总和
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果能分成两个子集,则相当于背包的重量为sum/2,而数组中的每个元素不可能为小数,所以判断
	if sum%2 != 0 {
		return false
	}
	packageWeight := sum / 2
	// 初始化动态规划的dp数组
	dp := make([]int, packageWeight+1)
	// 动态规划遍历-商品
	for _, num := range nums {
		// 动态规划遍历-背包
		for j := packageWeight; j >= num; j-- {
			// 为什么这里是max;因为目标是要把背包装满,才能证明可以刚好可以进行分割
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[packageWeight] == packageWeight
}

func canPartitionV2(nums []int) bool {
	// 计算出平均值
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 判断sum是否可以被2整除,如果不能被2整除;证明无法拆分成两个子集
	if sum%2 != 0 {
		return false
	}
	packageWeight := sum / 2
	// 定义动态规划一维数组
	dp := make([]int, packageWeight+1)
	// 初始化一维数组: dp[0] = 0
	// 动态规划一维遍历: 先遍历商品
	for _, num := range nums {
		for j := packageWeight; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[packageWeight] == packageWeight
}
